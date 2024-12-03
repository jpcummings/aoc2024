package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	nsafe := 0
	ndampersafe := 0
	for ( scanner.Scan() ) {
		line := scanner.Text()
		rep := strings.Split(line," ")
		var report []int
		for _, lvl := range rep {
			ilvl, _ := strconv.Atoi(lvl)
			report = append(report, ilvl)
		}
		if isSafe(report) {
			nsafe += 1
			ndampersafe += 1
		} else {
			//remove elements and try again
			for i, _ := range report {
				repcopy := append([]int(nil), report...)
				trep := remove(repcopy, i)
				if isSafe(trep) {
					ndampersafe += 1
					break
				}
			}
		}
	}
	// part 1
	fmt.Println(nsafe)

	// part 2
	fmt.Println(ndampersafe)

}

func isSafe(rep []int) bool {
	return isMonotonic(rep) && isGradual(rep)
}

func isGradual(rep []int) bool {
	var steps []int
	for i := 0; i < (len(rep)-1); i++ {
		steps = append(steps, rep[i+1]-rep[i])
	}
	for _,s := range steps {
		if s > 3 || s < -3 {
			return false
		}
	}
	return true
}

func isMonotonic(rep []int) bool {
	var dirs []int
	for i := 0; i < (len(rep)-1); i++ {
		dirs = append(dirs, direction(rep[i],rep[i+1]))
	}
	sum := 0
	for _,d := range dirs {
		sum += d
	}
	return int(math.Abs(float64(sum))) == len(dirs)
}

func direction(x,y int) int {
	dir := y-x
	if dir != 0 {
		dir = (y-x)/int(math.Abs(float64(y-x)))
	}
	return dir
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}