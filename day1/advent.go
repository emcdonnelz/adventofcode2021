package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ingestFile(fileName string) []string {
	contents, err := ioutil.ReadFile(fileName)
	check(err)
	values := strings.Split(string(contents), "\n")
	values = values[:len(values)-1]
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
	numberOfIncreases := 0

	previous := 0
	for index, value := range values {
		if index > 1 {
			first := convertStringToInt(values[index-2])
			second := convertStringToInt(values[index-1])
			third := convertStringToInt(value)
			sum := first + second + third
			if previous != 0 {
				if sum > previous {
					numberOfIncreases += 1
				}
			}
			previous = sum
		}
	}

	fmt.Println("Number Of Increases: ", numberOfIncreases)
}
