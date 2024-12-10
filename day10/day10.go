package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"strconv"
	"slices"
)

type Point struct {
	i, j int
}

type Cell struct {
	height int
	loc Point
}

var ninesfound []Cell

func main() {


	topomap := ReadTopomap(os.Stdin)
	// fmt.Println(topomap)
	fmt.Println(ScoreTopoMap(topomap))

}

func ScoreTopoMap(tmap [][]Cell) (int, int) {
	var score, score2  int
	for _, row := range tmap {
		for _, cell := range row {
			if cell.height == 0 {
				// found a trailhead, let's score it
				// fmt.Println("Trailhead: ", cell)
				score += ScoreTrailHeadUniq(tmap, cell)
				score2 +=ScoreTrailHead(tmap, cell)
				ResetNines()
			}
		}
	}
	return score, score2
}

func ScoreTrailHeadUniq(tmap [][]Cell, cell Cell) int {
	score := CountNinesUniq(tmap, cell)
	// fmt.Println("Score of ", cell, " is ", score, "\n\n")
	return score
}

func ScoreTrailHead(tmap [][]Cell, cell Cell) int {
	score := CountNines(tmap, cell)
	// fmt.Println("Score of ", cell, " is ", score, "\n\n")
	return score
}

func CountNinesUniq(tmap [][]Cell, cell Cell) int {
	var tot int
	neighbors := Neighbors(tmap, cell)
	// fmt.Println(neighbors)
	if cell.height == 9 && !slices.Contains(ninesfound, cell) {  // && !cell.found 
		// cell.found = true
		ninesfound = append(ninesfound, cell)
		return 1
	} else {
		for _, n := range neighbors {
			tot += CountNinesUniq(tmap, n)
		}
	}
	// fmt.Println("tot: ", tot)
	return tot
}

func CountNines(tmap [][]Cell, cell Cell) int {
	var tot int
	neighbors := Neighbors(tmap, cell)
	// fmt.Println(neighbors)
	if cell.height == 9  {  
		return 1
	} else {
		for _, n := range neighbors {
			tot += CountNines(tmap, n)
		}
	}
	// fmt.Println("tot: ", tot)
	return tot
}

func Neighbors(tmap [][]Cell, cell Cell) []Cell {
	var neighbors []Cell
	imax := len(tmap)-1
	jmax := len(tmap[0])-1
	if cell.loc.i-1 >= 0 {
		// fmt.Println("up", cell)
		n := tmap[cell.loc.i-1][cell.loc.j]
		if n.height == cell.height+1 {
			neighbors = append(neighbors, n)
		}
	}
	if cell.loc.i+1 <= imax {
		// fmt.Println("down", cell)
		n := tmap[cell.loc.i+1][cell.loc.j]
		if n.height == cell.height+1 {
			neighbors = append(neighbors, n)
		}
	}
	if cell.loc.j-1 >= 0 {
		// fmt.Println("left", cell)
		n := tmap[cell.loc.i][cell.loc.j-1]
		if n.height == cell.height+1 {
			neighbors = append(neighbors, n)
		}
	}
	if cell.loc.j+1 <= jmax {
		// fmt.Println("right", cell)
		n := tmap[cell.loc.i][cell.loc.j+1]
		if n.height == cell.height+1 {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func ReadTopomap(in io.Reader) [][]Cell {
	var topomap [][]Cell

	scanner := bufio.NewScanner(in)

	var i, j int
	for scanner.Scan() {
		var scanline []Cell
		line := scanner.Text()
		j = 0
		for _, c := range line {
			height,_ := strconv.Atoi(string(c))
			cell := Cell{height, Point{i,j}}
			scanline = append(scanline, cell)
			j++
		}
		topomap = append(topomap, scanline)
		i++
	}
	// fmt.Println(topomap)
	return topomap
}

func ResetNines() {
	ninesfound = nil
}
