package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"slices"
)

type Pair struct {
	first int
	second int
}

func main() {

	var section string= "rules"
	var rules []Pair
	var manuals [][]int
	var sum int = 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			section = "manuals"
			continue
		}
		if section == "rules" {
			pages := strings.Split(line, "|")
			first, _ := strconv.Atoi(pages[0])
			second, _ := strconv.Atoi(pages[1])
			rules = append(rules, Pair{first, second})
		} else if section == "manuals" {
			man := strings.Split(line, ",")
			var manual []int
			for _, p := range man {
				page, _ := strconv.Atoi(p)
				manual = append(manual, page)
			}
			manuals = append(manuals, manual)
		} else {
			fmt.Println("This can't happen")
		}
	}
	// fmt.Println(rules)
	// fmt.Println(manuals)

	for _, man := range manuals {
		if GoodManual(man,rules) {
			mid := FindMiddlePage(man)
			sum += mid
		}
	}
	fmt.Println(sum)

	sum = 0
	for _, man := range manuals {
		if !GoodManual(man,rules) {
			// FixManuals goes through list of rules moving pages to satisfy each rule.
			// Later rules may break earlier fixes, so repeat until page order is good.
			for !GoodManual(man,rules) {
				man = FixManual(man,rules)
			}
			mid := FindMiddlePage(man)
			sum += mid
		}
	}
	fmt.Println(sum)
}

func FindMiddlePage(man []int) int {
	return man[(len(man)-1)/2]
}

func GoodManual(manual []int, rules []Pair) bool {

	ret := true

	for _, rule := range rules {
		// are the pages in the manual?
		if PagesInMan(manual, rule) {
			// are they in the wrong order?
			if slices.Index(manual, rule.first) > slices.Index(manual, rule.second) {
				// wrong order!
				ret = false
			}
		} else {
			// fmt.Println(rule, " not in ", manual)
		}
	}
	return ret
}

func PagesInMan(man []int, rule Pair) bool {
	if slices.Contains(man, rule.first) && slices.Contains(man, rule.second) {
		return true
	} else {
		return false
	}
}

func FixManual(manual []int, rules []Pair) []int {

	for _, rule := range rules {
		// are the pages in the manual?
		if PagesInMan(manual, rule) {
			// are they in the wrong order?
			ifirst := slices.Index(manual, rule.first)
			isecond := slices.Index(manual, rule.second)
			if ifirst > isecond {
				//fmt.Println("Correcting rule: ", rule.first,"|",rule.second)
				// wrong order! fix it.
				manual = slices.Delete(manual, isecond, isecond+1)
				manual = slices.Insert(manual, ifirst, rule.second)
			}
		} else {
			// fmt.Println(rule, " not in ", manual)
		}
	}
	return manual
}

