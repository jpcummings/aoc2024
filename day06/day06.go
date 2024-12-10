package main

import (
	"bufio"
	"os"
	"fmt"
	"slices"
)

type Point struct {
	i int
	j int
}

type Guard struct {
	loc Point
	dir Point
}

type Visit struct {
	loc Point
	dir Point
}

func main() {
	var i int
	var grid [][]string
	var guard Guard

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for j, char  := range line {
			if string(char) == "^" {
				// fmt.Println("Guard at: ", i, ",", j)
				guard = Guard{Point{i,j}, Point{-1,0}}
				row = append(row, "X")
			} else {
				row = append(row, string(char))
			}
		}
		grid = append(grid, row)
		i++
	}
	//PrintMap(grid, guard)

	// p1
	var gone bool = false
	var visits []Visit
	guardstart := guard
	//fmt.Println(grid[guard.loc.i][guard.loc.j])
	for grid, guard, visits, gone, _ = walk(grid, guard, visits); gone == false; {
		guard = turn(guard)
		grid, guard, visits, gone, _ = walk(grid, guard, visits)
		//PrintMap(grid, guard)
	}
	cvisits := Compact(visits)
	fmt.Println(len(cvisits))

	// p2 - this is slooow.  rather than try blocks in _every_ spot, this should be rewritten to put blocks only in the path of the guard.
	// ... but it works, and I'm movin' on to day 7!
	nloops := 0
	loop := false
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] != "#" {
				// fmt.Println("trying block at ", i," ",  j)
				grid[i][j] = "#"
				guard = guardstart
				visits =visits[0:0]
				//PrintMap(grid, guard)
				// fmt.Println(visits)
				for grid, guard, visits, gone, loop = walk(grid, guard, visits); gone == false && loop == false; {
					guard = turn(guard)
					grid, guard, visits, gone, loop = walk(grid, guard, visits)
				}
				if loop == true {
					nloops++
				}
				grid[i][j] = "."
				// PrintMap(grid, guard)
				// fmt.Println(visits)
			}
		}
	}
	fmt.Println(nloops)
}

func PrintMap(grid [][]string, guard Guard) {
	for i, row := range grid {
		for j, el := range row {
			if i == guard.loc.i&& j == guard.loc.j {
				fmt.Print("g")
			} else {
				fmt.Print(el)
			}
		}
		fmt.Println()
	}
}

func walk(grid [][]string, guard Guard, visits []Visit) ([][]string, Guard, []Visit, bool, bool) {
	// fmt.Println("start: ",guard.loc)
	gone := false
	loop := false
	for {
		step := Point{guard.loc.i+guard.dir.i, guard.loc.j+guard.dir.j}
		if inbounds(grid, step) {
			if grid[step.i][step.j] == "#" {
				break
			}
			//grid[step.i][step.j] = "X"
			v := Visit{step,sub(step,guard.loc)}
			if Contains(visits, v) {
				loop = true
				return grid, guard, visits, gone, loop
			}
			visits = append(visits, v)
			guard.loc = step
		} else {
			gone = true
			break
		}
	}
	return grid, guard, visits, gone, loop
}

func turn(guard Guard) Guard {
	// 90 deg right
	i := guard.dir.j
	j := -guard.dir.i
	guard.dir = Point{i,j}
	return guard
}

func inbounds(grid [][]string, p Point) bool {
	if p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[0]) {
		return false
	}
	return true
}

func sub(x, y Point) Point {
	return Point{x.i-y.i,x.j-y.j}
}

func add(x, y Point) Point {
	return Point{x.i+y.i,x.j+y.j}
}

func Compact(vs []Visit) []Visit {
	for i:=0; i<len(vs)-2; i++ {
		for j:=i+1; j<len(vs)-1; j++ {
			if vs[i].loc == vs[j].loc {
				vs = slices.Delete(vs, j, j+1)
			}
		}
	}
	return vs
}

func Loop(vs []Visit) bool {
	for i:=0; i<len(vs)-2; i++ {
		for j:=i+1; j<len(vs)-1; j++ {
			if vs[i].loc == vs[j].loc && vs[i].dir == vs[j].dir{
				return true
			}
		}
	}
	return false
}

func Contains(vs []Visit, v Visit) bool {
	for _, vl := range vs {
		if vl.loc == v.loc && vl.dir == v.dir{
			return true
		}
	}
	return false
}
