package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Tracker struct {
	Values []int
	Extra  []int
}

func (t *Tracker) CalcFirstSevenDays() int {
	numberOfValues := len(t.Extra)
	for f := 0; f < numberOfValues; f++ {
		switch t.Extra[f] {
		case 0:
			t.Extra[f] = 6
			t.Extra = append(t.Extra, 9)
			numberOfValues += 1
		default:
			t.Extra[f] -= 1
		}
	}
	return numberOfValues
}

func (t *Tracker) ExecuteDays(days int) int {
	total := 0 // Number of initial fishes

	if days <= 7 {
		total += t.CalcFirstSevenDays()
	} else {
		days += 1              //assuming caller started at day 0
		total += len(t.Values) // add all initial fish from input
		for f := 0; f < len(t.Values); f++ {
			initial := t.Values[f]
			daysLeftAtFirstBaby := days - initial - 1
			r, extraDays := daysLeftAtFirstBaby/7, daysLeftAtFirstBaby%7
			total += r //first generation babies
			total += 1 // initial birth at daysLeftAtFirstBaby
			rounds := []int{}
			counter := 4
			skip := 4
			skip_counter := 4
			for i := 1; i <= r; i++ {
				if i == skip {
					// I don't understand this skip counter (+4 or +5) but it works
					skip += skip_counter
					switch skip_counter {
					case 4:
						skip_counter = 5
					case 5:
						skip_counter = 4
					}
				} else {
					e := 0
					order := 0
					switch counter {
					case 4:
						order = 2
					case 2:
						order = 4
					case 0:
						order = 6
					case 5:
						order = 1
					case 3:
						order = 3
					case 1:
						order = 5
					}
					if extraDays >= order {
						e += 1
					}
					rounds = append(rounds, r-i+e)
					switch counter {
					case 4:
						counter = 2
					case 2:
						counter = 0
					case 0:
						counter = 5
					case 5:
						counter = 3
					case 3:
						counter = 1
					case 1:
						counter = 6
					case 6:
						counter = 4
					}
				}
			}

			firstGen := []int{}
			if len(rounds) > 0 {
				for i := 1; i <= rounds[0]; i++ {
					firstGen = append(firstGen, 1)
				}
				total += calculate(rounds, firstGen)
			}
		}
	}
	return total
}

func calculate(rounds []int, lastGen []int) int {
	var total, previous int = 0, 0
	thisGen := []int{}
	numberOfRounds := rounds[0]
	rounds = rounds[1:]
	for i := 0; i < numberOfRounds; i++ {
		new := previous + lastGen[i]
		previous = previous + lastGen[i]
		total += new
		thisGen = append(thisGen, new)
	}
	if len(rounds) <= 0 {
		return total
	}
	return total + calculate(rounds, thisGen)
}

func ingestFile(fileName string) []int {
	contents, err := ioutil.ReadFile(fileName)
	check(err)
	values := strings.Split(string(contents), "\n")
	values = strings.Split(string(values[0]), ",")
	valuesInt := []int{}
	for _, value := range values {
		valuesInt = append(valuesInt, convertStringToInt(value))
	}
	return valuesInt
}

func convertStringToInt(strValue string) int {
	intValue, err := strconv.Atoi(strValue)
	check(err)
	return intValue
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	values := ingestFile("input")
	tracker := Tracker{
		Values: values,
	}
	days := 255
	total := tracker.ExecuteDays(days)
	fmt.Println("Number of fishes: ", total)
}
