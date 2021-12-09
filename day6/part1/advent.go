package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Fish struct {
	Counter int
}

type Tracker struct {
	Fishes []Fish
}

func (t *Tracker) GetNumberOfFishes() int {
	return len(t.Fishes)
}

func (t *Tracker) ExecuteDayEnd() {
	for i := 0; i < len(t.Fishes); i++ {
		switch t.Fishes[i].Counter {
		case 0:
			t.Fishes[i].Counter = 6
			newFish := Fish{
				Counter: 9,
			}
			t.Fishes = append(t.Fishes, newFish)
		default:
			t.Fishes[i].Counter -= 1
		}
	}
}

func (t *Tracker) CreateInitialFish(values []string) {
	for i := 0; i < len(values); i++ {
		fish := Fish{
			Counter: convertStringToInt(values[i]),
		}
		t.Fishes = append(t.Fishes, fish)
	}
}

func ingestFile(fileName string) []string {
	contents, err := ioutil.ReadFile(fileName)
	check(err)
	values := strings.Split(string(contents), "\n")
	values = strings.Split(string(values[0]), ",")
	return values
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertStringToInt(strValue string) int {
	intValue, err := strconv.Atoi(strValue)
	check(err)
	return intValue
}

func main() {
	values := ingestFile("input")
	tracker := Tracker{
		Fishes: []Fish{},
	}
	tracker.CreateInitialFish(values)
	days := 80
	fmt.Print("Calculating day:")
	for i := 0; i < days; i++ {
		fmt.Print(" ", i)
		tracker.ExecuteDayEnd()
	}
	fmt.Println(" ")
	fmt.Println("Number of fishes: ", tracker.GetNumberOfFishes())
}
