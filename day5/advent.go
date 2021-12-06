package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	X       int
	Y       int
	Overlap int
}

type Line struct {
	Points []Point
}

type Board struct {
	Lines      []Line
	Duplicates int
}

func (board *Board) AddNewLine(fullCoordinates string) {
	endpoints := strings.Split(string(fullCoordinates), " -> ")
	startPoint := strings.Split(string(endpoints[0]), ",")
	endPoint := strings.Split(string(endpoints[1]), ",")
	line := Line{
		Points: []Point{},
	}
	line.AddPoints(startPoint, endPoint)
	board.Lines = append(board.Lines, line)
	board.CheckNewLineOverlaps()
}

func (board *Board) CheckNewLineOverlaps() {
	numberOfLines := len(board.Lines)
	if numberOfLines > 1 {
		line := board.Lines[numberOfLines-1]
		for newP := 0; newP < len(line.Points); newP++ {
			for oldL := 0; oldL < numberOfLines-1; oldL++ {
				for oldP := 0; oldP < len(board.Lines[oldL].Points); oldP++ {
					if line.Points[newP].EqualToPoint(board.Lines[oldL].Points[oldP]) {
						line.Points[newP].Overlap = 1
						if board.Lines[oldL].Points[oldP].Overlap == 0 {
							board.Lines[oldL].Points[oldP].Overlap = 1
							board.Duplicates += 1
						}
					}
				}
			}
		}
	}
}

func (line *Line) AddPoints(startPoint []string, endPoint []string) {
	startX := convertStringToInt(startPoint[0])
	startY := convertStringToInt(startPoint[1])
	endX := convertStringToInt(endPoint[0])
	endY := convertStringToInt(endPoint[1])

	direction := getDirection(startX, startY, endX, endY)

	switch direction {
	case "single-point":
		line.Points = append(line.Points, createNewPoint(startX, startY))
	case "vertical":
		line.AddHorizonOrVirtPoints(startY, endY, startX, false)
	case "horizontal":
		line.AddHorizonOrVirtPoints(startX, endX, startY, true)
	case "diagonal":
		line.AddDiagonalPoints(startX, startY, endX, endY)
	}
}

func (line *Line) AddDiagonalPoints(startX int, startY int, endX int, endY int) {
	if (startX < endX) && (startY < endY) {
		// Down & Right
		for (startX <= endX) && (startY <= endY) {
			line.Points = append(line.Points, createNewPoint(startX, startY))
			startX += 1
			startY += 1
		}
	} else if (startX > endX) && (startY < endY) {
		// Down & Left
		for (startX >= endX) && (startY <= endY) {
			line.Points = append(line.Points, createNewPoint(startX, startY))
			startX -= 1
			startY += 1
		}
	} else if (startX < endX) && (startY > endY) {
		// Up & Right
		for (startX <= endX) && (startY >= endY) {
			line.Points = append(line.Points, createNewPoint(startX, startY))
			startX += 1
			startY -= 1
		}
	} else if (startX > endX) && (startY > endY) {
		// Up & Left
		for (startX >= endX) && (startY >= endY) {
			line.Points = append(line.Points, createNewPoint(startX, startY))
			startX -= 1
			startY -= 1
		}
	} else {
		fmt.Println("Oops!")
	}
}

func (line *Line) AddHorizonOrVirtPoints(start int, end int, other int, isHorizon bool) {
	pointValue := start
	loopStart := start
	loopStop := end
	if start > end {
		loopStart = end
		loopStop = start
	}
	for loopStart <= loopStop {
		newPoint := Point{}
		if isHorizon {
			newPoint = createNewPoint(pointValue, other)
		} else {
			newPoint = createNewPoint(other, pointValue)
		}
		line.Points = append(line.Points, newPoint)
		loopStart += 1
		if start > end {
			pointValue = pointValue - 1
		} else {
			pointValue = pointValue + 1
		}
	}
}

func (point *Point) EqualToPoint(otherPoint Point) bool {
	return (point.X == otherPoint.X) && (point.Y == otherPoint.Y)
}

func getDirection(startX int, startY int, endX int, endY int) string {
	direction := "diagonal"
	if (startX == endX) && (startY == endY) {
		direction = "single-point"
	} else if startX == endX {
		direction = "vertical"
	} else if startY == endY {
		direction = "horizontal"
	}
	return direction
}

func createNewPoint(x int, y int) Point {
	point := Point{
		X:       x,
		Y:       y,
		Overlap: 0,
	}
	return point
}

func createBoard(boardContents []string) Board {
	board := Board{
		Lines:      []Line{},
		Duplicates: 0,
	}
	for l := 0; l < len(boardContents); l++ {
		board.AddNewLine(boardContents[l])
	}
	return board
}

func convertStringToInt(strValue string) int {
	intValue, err := strconv.Atoi(strValue)
	check(err)
	return intValue
}

func ingestFile(fileName string) []string {
	contents, err := ioutil.ReadFile(fileName)
	check(err)
	coordinates := strings.Split(string(contents), "\n")
	coordinates = coordinates[:len(coordinates)-1]
	return coordinates
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	contents := ingestFile("input")
	board := createBoard(contents)

	fmt.Println("Board Duplicates: ", board.Duplicates)
}
