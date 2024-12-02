package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
	"math"
)

func main() {
	var left,right []int
	var i int
	scanner := bufio.NewScanner(os.Stdin)

	for ( scanner.Scan() ) {
		line := scanner.Text()
		tmp := strings.Fields(line)
		i, _ = strconv.Atoi(tmp[0])
		left = append(left,i)
		i, _ = strconv.Atoi(tmp[1])
		right = append(right,i)
	}

	// part 1
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))
	tot := 0
	for i, _ := range left {
		tot += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println(tot)

	// part 2
	num := 0  // assume 0 is not in list?
	nright := 0
	tot = 0
	for i, _ := range left {
	// find next val
		if left[i] != num {
			num = left[i]
			// count nums in right list
			nright = 0
			for j, _ := range right {
				if right[j] == num {
					nright += 1
				}
				if right[j] > num {
					break
				}
			}
		}
		tot += num*nright
	}
	fmt.Println(tot)

}

