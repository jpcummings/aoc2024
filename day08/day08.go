package main

import (
	"bufio"
	"os"
	"fmt"
)

type Point struct {
	i int
	j int
}

func main() {
	var grid [][]string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var row []string
		line := scanner.Text()
		for _,c := range line {
			row = append(row, string(c))
		}
		grid = append(grid, row)
	}
	antenna := FindAntenna(grid)
	imax := len(grid)
	jmax := len(grid[0])
	nodes := FindNodes(antenna, imax, jmax)
	fmt.Println(CountNodes(nodes))
	nodes = FindAllNodes(antenna, imax, jmax)
	fmt.Println(CountNodes(nodes))
}

func PrintMap(m [][]string, nodes []Point) {
	for _, p := range nodes {
		if m[p.i][p.j] == "." {
			m[p.i][p.j] = "#"
		}
	}
	for _, row := range m {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func FindAntenna(m [][]string) map[string][]Point {
	ant := make(map[string][]Point)
	for i, row := range m {
		for j, el := range row {
			if el != "." {
				ant[el] = append(ant[el], Point{i,j})
			}
		}
	}
	return ant
}

func FindNodes(antennas map[string][]Point, imax, jmax int) []Point{
	var node []Point

	for _, antenna := range antennas {
		for i, a1 := range antenna {
			for j, a2 := range antenna {
				if i > j {
					del := sub(a2, a1)
					n1 := sub(a1, del)
					n2 := add(a2, del)
					if inbounds(n2, imax, jmax) {
						node = append(node, n2)
					}
					if inbounds(n1, imax, jmax) {
						node = append(node, n1)
					}
				}
			}
		}
	}
	return node
}

func FindAllNodes(antennas map[string][]Point, imax, jmax int) []Point{
	var node []Point

	for _, antenna := range antennas {
		for i, a1 := range antenna {
			for j, a2 := range antenna {
				if i > j {
					node = append(node, a1)
					node = append(node, a2)
					// fmt.Println(a1, " and ", a2)
					del := sub(a2, a1)
					del = Point{del.i/gcd(del.i,del.j), del.j/gcd(del.i,del.j)}
					n1 := sub(a1, del)
					for inbounds(n1, imax, jmax) {
						node = append(node, n1)
						n1 = sub(n1, del)
					}
					n2 := add(a2, del)
					for inbounds(n2, imax, jmax) {
						node = append(node, n2)
						n2 = add(n2, del)
					}
				}
			}
		}
	}
	return node
}

func CountNodes(points []Point) int {
	m := make(map[Point]int)
	for _,p := range points {
		m[p]++
	}
	return len(m)
}

func inbounds(p Point, imax, jmax int) bool {
	if p.i > -1 && p.i < imax && p.j > -1 && p.j < jmax {
		return true
	} else {
		return false
	}
}


func add(p1, p2 Point) Point {
	return Point{p1.i+p2.i, p1.j+p2.j}
}

func sub(p1, p2 Point) Point {
	return Point{p1.i-p2.i, p1.j-p2.j}
}

func gcd(x, y int) int {  // Eulers method!  Hadn't used this in a while...
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}
