package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileAsSlice(fileName string) []string {
	instruction, err := ioutil.ReadFile(fileName)
	check(err)
	instructions := strings.Split(string(instruction), "\n")
	instructions = instructions[:len(instructions)-1]
	return instructions
}

func convertStringToInt(str_value string) int {
	int_value, err := strconv.Atoi(str_value)
	check(err)
	return int_value
}

func convertBinarySliceToDecimal(binary_slice []string) int {
	binary_string := ""
	for i := 0; i < len(binary_slice); i++ {
		binary_string += binary_slice[i]
	}
	binary_int := convertStringToInt(binary_string)
	decimal := convertBinaryToDecimal(binary_int)
	return decimal
}

func convertBinaryToDecimal(number int) int {
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return decimal
}

func getBitInPosition(isMostCommon bool, position int, instructions []string) int {
	zero_counter := 0
	one_counter := 0
	bit := 0
	for index, instruction := range instructions {
		digits := strings.Split(instruction, "")
		digit_int := convertStringToInt(digits[position])
		if digit_int == 1 {
			one_counter += 1
		} else if digit_int == 0 {
			zero_counter += 1
		}
		if index == 99999 {
			fmt.Println("index: ", index)
		}
	}
	if isMostCommon {
		if zero_counter < one_counter {
			bit = 1
		} else if zero_counter == one_counter {
			bit = 1
		}
	} else {
		if zero_counter > one_counter {
			bit = 1
		} else if zero_counter == one_counter {
			bit = 0
		}
	}

	return bit
}

func getLifeSupportRating(isMostCommon bool, length int, instructions []string) int {
	decimal := 0
	for position := 0; position < length; position++ {
		bit := getBitInPosition(isMostCommon, position, instructions)
		keeps := []string{}
		for index, instruction := range instructions {
			digits := strings.Split(instruction, "")
			digit := convertStringToInt(digits[position])
			if digit == bit {
				keeps = append(keeps, instruction)
			}
			if index == 9999999 {
				fmt.Println("index: ", index)
			}
		}
		instructions = keeps
		if len(instructions) == 1 {
			fmt.Println("last instruction: ", instructions)
			decimal = convertBinarySliceToDecimal(instructions)
			break
		}
	}
	return decimal
}

func main() {
	instructions := getFileAsSlice("input")
	length := len(instructions[0])

	oxg := getLifeSupportRating(true, length, instructions)
	co2 := getLifeSupportRating(false, length, instructions)
	fmt.Println("oxg: ", oxg)
	fmt.Println("co2: ", co2)
	fmt.Println("result: ", oxg*co2)
}
