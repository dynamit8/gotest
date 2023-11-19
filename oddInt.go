package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getOddInt(intList []int) int {
	var checked = []int{}
	slices.Sort(intList)
	fmt.Printf("intList: %v\n", intList)
	for _, value := range intList {
		if slices.Index(checked, value) == -1 && len(checked) != 0 {
			if len(checked)%2 == 1 {
				return checked[0]
			} else {
				fmt.Printf("Old even value: %v\n", checked)
				checked = []int{}
				checked = append(checked, value)
			}
		} else {
			checked = append(checked, value)
		}
	}
	return checked[0]
}

func getOddIntPromt() int {
	promt := "Input list of integers that only 1 number occurs odd number of times\nExample:\n1,2,55,2,3,5,11,1\n\n"
	input := getPromt(promt)
	numbers, convertResult := validateOddInput(input)
	if !convertResult {
		return -1
	}
	result := getOddInt(numbers)
	return result
}

func validateOddInput(input string) ([]int, bool) {
	substrings := strings.Split(input, ",")
	var numbers []int

	for _, str := range substrings {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return numbers, false
		}
		numbers = append(numbers, num)
	}
	return numbers, true
}
