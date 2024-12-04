package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	nxmas := 0
	var grid []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid,line)
	}


	for i, row := range grid {
		for j, _ := range row {
			nxmas += findXMAS(grid,i,j)
		}
	}
	fmt.Println(nxmas)

	nxmas = 0
	for i, row := range grid {
		for j, _ := range row {
			nxmas += findX_MAS(grid,i,j)
		}
	}
	fmt.Println(nxmas)
}

func findX_MAS(grid []string, i, j int) int {

	nxmas := 0

	if string(grid[i][j]) == "A" && inbounds(grid,i,j) {
		if string(grid[i-1][j-1]) == "M" && string(grid[i+1][j+1]) == "S" && string(grid[i+1][j-1]) == "M" && string(grid[i-1][j+1]) == "S"{
			nxmas++
		}
		if string(grid[i+1][j-1]) == "M" && string(grid[i-1][j+1]) == "S" && string(grid[i+1][j+1]) == "M" && string(grid[i-1][j-1]) == "S"{
			nxmas++
		}
		if string(grid[i-1][j-1]) == "M" && string(grid[i+1][j+1]) == "S" && string(grid[i-1][j+1]) == "M" && string(grid[i+1][j-1]) == "S"{
			nxmas++
		}
		if string(grid[i-1][j+1]) == "M" && string(grid[i+1][j-1]) == "S" && string(grid[i+1][j+1]) == "M" && string(grid[i-1][j-1]) == "S"{
			nxmas++
		}
	}
	return nxmas
}

func inbounds(grid []string, i, j int) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	
	if i>0 && i+1<nrows && j>0 && j+1<ncols {
		return true
	}
	return false
}

func findXMAS(grid []string, i, j int) int {

	nxmas := 0
	nrows := len(grid)
	ncols := len(grid[0])

	if string(grid[i][j]) == "X" {
		// look right
		if (j+3)<ncols && string(grid[i][j+1]) == "M" && string(grid[i][j+2]) == "A" && string(grid[i][j+3]) == "S" {
			nxmas++
		}
		// look left
		if (j-3)>-1 && string(grid[i][j-1]) == "M" && string(grid[i][j-2]) == "A" && string(grid[i][j-3]) == "S" {
			nxmas++
		}
		// look up
		if (i-3)>-1 && string(grid[i-1][j]) == "M" && string(grid[i-2][j]) == "A" && string(grid[i-3][j]) == "S" {
			nxmas++
		}
		// look down
		if (i+3)<nrows && string(grid[i+1][j]) == "M" && string(grid[i+2][j]) == "A" && string(grid[i+3][j]) == "S" {
			nxmas++
		}
		// look upright
		if (i-3)>-1 && (j+3)<ncols && string(grid[i-1][j+1]) == "M" && string(grid[i-2][j+2]) == "A" && string(grid[i-3][j+3]) == "S" {
			nxmas++
		}
		// look downleft
		if (i+3)<nrows && (j-3)>-1 && string(grid[i+1][j-1]) == "M" && string(grid[i+2][j-2]) == "A" && string(grid[i+3][j-3]) == "S" {
			nxmas++
		}
		// look uplef
		if (i-3)>-1 && (j-3)>-1 && string(grid[i-1][j-1]) == "M" && string(grid[i-2][j-2]) == "A" && string(grid[i-3][j-3]) == "S" {
			nxmas++
		}
		// look downright
		if (i+3)<nrows && (j+3)<ncols && string(grid[i+1][j+1]) == "M" && string(grid[i+2][j+2]) == "A" && string(grid[i+3][j+3]) == "S" {
			nxmas++
		}
	}

	return nxmas
}