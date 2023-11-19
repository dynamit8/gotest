package main

import (
	"fmt"
	"slices"
	"strings"
)

var VALID_EYES = []string{":", ";"}
var VALID_NOSE = []string{"-", "~"}
var VALID_MOUTH = []string{")", "D"}
var VALID_SMILEY = getValidSmiley()

func getValidSmiley() []string {
	var fullresult string
	var result string
	var fullSmiley = []string{}
	var smiley = []string{}
	for _, eye := range VALID_EYES {
		for _, nose := range VALID_NOSE {
			for _, mouth := range VALID_MOUTH {
				fullresult = eye + nose + mouth
				result = eye + mouth
				if slices.Index(smiley, result) == -1 {
					smiley = append(smiley, result)
				}
				fullSmiley = append(fullSmiley, fullresult)
			}
		}
	}
	allSmiley := append(fullSmiley, smiley...)
	return allSmiley
}

func checkValidSmile(inputList []string) int {
	var result int = 0
	if len(inputList) == 0 {
		return result
	}
	for _, face := range inputList {
		if len(face) < 2 || len(face) > 3 {
			continue
		}
		if isContain(face, VALID_SMILEY) {
			result++
		}
	}
	return result
}

func isContain(character string, strList []string) bool {
	for _, value := range strList {
		fmt.Printf("value and character: %v - %v\n", value, character)
		if value == character {
			return true
		}
	}
	return false
}

func getSmileyFacesPromt() int {
	promt := "Input list of smiley faces\nExample:\n:) ;( ;} :-D\n\n"
	input := getPromt(promt)
	stringList := convertToStringSlice(input)
	result := checkValidSmile(stringList)
	return result
}

func convertToStringSlice(input string) []string {
	substrings := strings.Split(input, " ")
	var smileys []string
	smileys = append(smileys, substrings...)
	return smileys
}
