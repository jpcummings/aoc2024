package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"slices"
)

func main() {

	var val int
	var nums []int
	var totalcal, totalcal2 int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		val, nums = parseLine(line)
		testvals := eval(nums)
		testvalscat := evalcat(nums)
		if slices.Contains(testvals,val) {
			totalcal += val
		} else if slices.Contains(testvalscat,val) {
			totalcal2 += val
		}
	}
	fmt.Println(totalcal)
	fmt.Println(totalcal + totalcal2)
}

func eval(nums []int) []int {
	var ret []int
	if len(nums) == 1 {
		return nums
	}
	vals := eval(nums[1:])
	for _, v := range vals {
		ret = append(ret, nums[0]+v)
		ret = append(ret, nums[0]*v)
	}
	return ret
}

func evalcat(nums []int) []int {
	var ret []int
	if len(nums) == 1 {
		return nums
	}
	vals := evalcat(nums[1:])
	for _, v := range vals {
		ret = append(ret, nums[0]+v)
		ret = append(ret, nums[0]*v)
		ret = append(ret, cat(v, nums[0]))
	}
	return ret
}

func cat(x, y int) int {
	ret, _ := strconv.Atoi(strconv.Itoa(x)+strconv.Itoa(y))
	return ret
}

func parseLine(line string) (int, []int) {

	var nums [] int

	p := strings.Split(line, ": ")
	val, _ := strconv.Atoi(p[0])
	snums := strings.Split(p[1], " ")
	for _, snum := range snums {
		num, _ := strconv.Atoi(snum)
		nums = append(nums, num)
	}
	slices.Reverse(nums)
	return val, nums
}