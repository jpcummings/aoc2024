package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var sum1, sum2 int
	var mul int = 1
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		reinstr := regexp.MustCompile(`(mul\(\d+,\d+\))|(don\'t\(\))|(do\(\))`)
		instructions := reinstr.FindAll([]byte(line), -1)
		for _,instr := range instructions {
			if string(instr) == "do()" {
				mul = 1
			} else if string(instr)  == "don't()" {
				mul = 0
			} else {
				renum := regexp.MustCompile(`\d+`)
				nums := renum.FindAll([]byte(instr), -1)
				num0, _ := strconv.Atoi(string(nums[0]))
				num1, _ := strconv.Atoi(string(nums[1]))
				sum1 += num0*num1
				sum2 += mul*num0*num1
			}
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}