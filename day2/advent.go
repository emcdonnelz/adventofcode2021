package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (p *Position) Forward(amount int) {
	p.Horizontal += amount
	p.Depth += (p.Aim * amount)
}

func (p *Position) Down(amount int) {
	p.Aim += amount
}

func (p *Position) Up(amount int) {
	p.Aim -= amount
}

func (p *Position) GetValue() int {
	return p.Horizontal * p.Depth
}

func ingestFile(fileName string) []string {
	contents, err := ioutil.ReadFile(fileName)
	check(err)
	instructions := strings.Split(string(contents), "\n")
	instructions = instructions[:len(instructions)-1]
	return instructions
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
	instructions := ingestFile("input")
	position := Position{
		Horizontal: 0,
		Depth:      0,
		Aim:        0,
	}
	for _, instruction := range instructions {
		direction := strings.Split(instruction, " ")
		switch direction[0] {
		case "forward":
			position.Forward(convertStringToInt(direction[1]))
		case "down":
			position.Down(convertStringToInt(direction[1]))
		case "up":
			position.Up(convertStringToInt(direction[1]))
		}
	}

	fmt.Println("Position: ", position)
	fmt.Println("position: ", position.GetValue())
}
