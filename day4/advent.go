package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Winner struct {
	BoardIndex int
	Points     int
}

func (w *Winner) GetWinnerPoints() int {
	return w.Points
}

type WinnerBoard struct {
	Winners        []Winner
	Full           bool
	NumberOfBoards int
}

func (wb *WinnerBoard) CheckIfWinnerBoardFull() {
	if len(wb.Winners) == wb.NumberOfBoards {
		wb.Full = true
	}
}

func (wb *WinnerBoard) DoesBoardIndexExist(boardIndex int) bool {
	doesBoardExist := false
	for i := 0; i < len(wb.Winners); i++ {
		if wb.Winners[i].BoardIndex == boardIndex {
			doesBoardExist = true
			break
		}
	}
	return doesBoardExist
}

func (wb *WinnerBoard) AddWinner(boardIndex int, points int) {
	wb.CheckIfWinnerBoardFull()
	if !(wb.Full) && !(wb.DoesBoardIndexExist(boardIndex)) {
		winner := Winner{
			BoardIndex: boardIndex,
			Points:     points,
		}
		wb.Winners = append(wb.Winners, winner)
	}
}

func (wb *WinnerBoard) GetFirstPlace() int {
	return wb.Winners[0].GetWinnerPoints()
}

func (wb *WinnerBoard) GetLastPlace() int {
	return wb.Winners[wb.NumberOfBoards-1].GetWinnerPoints()
}

type Row struct {
	Values []string
}

func (r *Row) GetValueAtPosition(position int) string {
	return r.Values[position]
}

func (r *Row) IsWinnerValueAtPosition(position int) bool {
	isWinner := false
	if r.Values[position] == "999" {
		isWinner = true
	}
	return isWinner
}

func (r *Row) GetLength() int {
	return len(r.Values)
}

func (r *Row) UpdateValueAtPosition(position int, newValue string) {
	r.Values[position] = newValue
}

func (r *Row) IsRowWinner() bool {
	isWinner := false
	winnerCount := len(r.Values)
	for i := 0; i < len(r.Values); i++ {
		if r.IsWinnerValueAtPosition(i) {
			winnerCount = winnerCount - 1
		}
	}
	if winnerCount == 0 {
		isWinner = true
	}
	return isWinner
}

func (r *Row) GetWinnerPoints() int {
	winnerPoints := 0
	for i := 0; i < len(r.Values); i++ {
		if !(r.IsWinnerValueAtPosition(i)) {
			rowValue := convertStringToInt(r.GetValueAtPosition(i))
			winnerPoints += rowValue
		}
	}
	return winnerPoints
}

type Board struct {
	Rows []Row
}

func (b *Board) AddRow(newRow []string) {
	row := Row{
		Values: newRow,
	}
	b.Rows = append(b.Rows, row)
}

func (b *Board) UpdateBoard(callingNumber int) {
	for rowIndex := 0; rowIndex < len(b.Rows); rowIndex++ {
		row := b.Rows[rowIndex]
		for valueIndex := 0; valueIndex < row.GetLength(); valueIndex++ {
			rowValue := row.GetValueAtPosition(valueIndex)
			rowInt := convertStringToInt(rowValue)
			if rowInt == callingNumber {
				row.UpdateValueAtPosition(valueIndex, "999")
			}
		}
	}
}

func (b *Board) IsHorizontalWinner() bool {
	isWinner := false
	for rowIndex := 0; rowIndex < len(b.Rows); rowIndex++ {
		if b.Rows[rowIndex].IsRowWinner() {
			isWinner = true
			break
		}
	}
	return isWinner
}

func (b *Board) IsVirticalWinner() bool {
	isWinner := false
	height := len(b.Rows)
	width := b.Rows[0].GetLength()
	for horizontalIndex := 0; horizontalIndex < width; horizontalIndex++ {
		winnerCount := height
		for virticalIndex := 0; virticalIndex < height; virticalIndex++ {
			if b.Rows[virticalIndex].IsWinnerValueAtPosition(horizontalIndex) {
				winnerCount = winnerCount - 1
			}
		}
		if winnerCount == 0 {
			isWinner = true
			break
		}
	}

	return isWinner
}

func (b *Board) IsBoardWinner() bool {
	isWinner := false

	if b.IsHorizontalWinner() {
		isWinner = true
	} else {
		if b.IsVirticalWinner() {
			isWinner = true
		}
	}
	return isWinner
}

func (b *Board) GetWinnerPoints(callingNumber int) int {
	winnerPoints := 0
	for rowIndex := 0; rowIndex < len(b.Rows); rowIndex++ {
		winnerPoints += b.Rows[rowIndex].GetWinnerPoints()
	}
	winnerPoints = winnerPoints * callingNumber
	return winnerPoints
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

func getFileAsSlice(fileName string) []string {
	instruction, err := ioutil.ReadFile(fileName)
	check(err)
	instructions := strings.Split(string(instruction), "\n")
	instructions = instructions[:len(instructions)-1]
	keeps := []string{}
	for index, instruction := range instructions {
		if instruction != "" {
			keeps = append(keeps, instructions[index])
		}
	}
	return keeps
}

func getRowValues(row string) []string {
	rowSlice := strings.Split(string(row), " ")
	keeps := []string{}
	for index, rowValue := range rowSlice {
		if rowValue != "" {
			keeps = append(keeps, rowSlice[index])
		}
	}
	return keeps
}

func splitBoards(boards []string) []Board {
	var allBoards = []Board{}
	for topRow := 0; topRow < len(boards); topRow += 5 {
		newBoard := Board{}
		for rowIndex := 0; rowIndex < 5; rowIndex++ {
			newBoard.AddRow(getRowValues(boards[topRow+rowIndex]))
		}
		allBoards = append(allBoards, newBoard)
	}
	return allBoards
}

func getCallingNumbers(allNumbers string) []int {
	var callingNumbers = []int{}
	numbersSlice := strings.Split(string(allNumbers), ",")

	for c := 0; c < len(numbersSlice); c++ {
		callInt := convertStringToInt(numbersSlice[c])
		callingNumbers = append(callingNumbers, callInt)
	}
	return callingNumbers
}

func main() {
	instructions := getFileAsSlice("input")

	callingNumbers := getCallingNumbers(instructions[0])
	allBoards := splitBoards(instructions[1:])

	winnerBoard := WinnerBoard{
		Winners:        []Winner{},
		Full:           false,
		NumberOfBoards: len(allBoards),
	}

	for call := 0; call < len(callingNumbers); call++ {
		for boardIndex := 0; boardIndex < len(allBoards); boardIndex++ {
			allBoards[boardIndex].UpdateBoard(callingNumbers[call])
			if allBoards[boardIndex].IsBoardWinner() {
				winnerPoints := allBoards[boardIndex].GetWinnerPoints(callingNumbers[call])
				winnerBoard.AddWinner(boardIndex, winnerPoints)
			}
		}
	}
	fmt.Println("First Place: ", winnerBoard.GetFirstPlace())
	fmt.Println("Last Place: ", winnerBoard.GetLastPlace())
	fmt.Println("WinnerBoard: ", winnerBoard)
}
