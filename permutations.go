package main

import (
	"fmt"
	"slices"
	"strings"
)

func getPermuteValidate(promt string, input string) string {
	if !validateEmptyStr(input) {
		panic("Please input non-empty text")
	}
	if !validateSpaceInString(input) {
		panic("Please input non-space text")
	}
	return input
}

func getPermutationPromt() []string {
	promt := "Please input to calculate permutations(Not empty, No space): "
	input := getPromt(promt)
	cleanStr := getPermuteValidate(promt, input)
	fmt.Printf("-- input: %v -- \n", cleanStr)
	permulationsList := getPermutation(cleanStr)
	return permulationsList
}

func getPermutation(input string) []string {
	var result = []string{}
	if len(input) == 1 {
		result = append(result, input)
		return result
	}

	temp := strings.Clone(input)
	for i := 0; i < len(input); i++ {
		var perms = []string{}
		head := temp[:1]
		tail := temp[1:]
		perms = getPermutation(tail)
		for _, perm := range perms {
			perm += head
			if slices.Index(result, perm) == -1 {
				result = append(result, perm)
			}
		}
		temp = rightShiftString(temp)
	}
	return result
}

func rightShiftString(s string) string {
	result := s[1:] + s[:1]
	return result
}
