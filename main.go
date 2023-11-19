package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Test Project")
	displayChoices()
	// result := getPermutations()
	// fmt.Println(result)

	// var iList = []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 1, 2, 2, 1, 3, 4, 5, 3, 3, 2, 2, 1}
	// odd := getOddInt(iList)
	// fmt.Println(odd)

	// allSmile := checkValidSmile()
	// fmt.Println(allSmile)
	// checkValidSmile()
}

func getInput(promt string, reader *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, error := reader.ReadString('\n')
	cleanStr := strings.TrimSpace(input)
	return cleanStr, error
}

func getPromt(str string) string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := getInput(str, reader)
	return input
}

func displayChoices() {
	input := getPromt("Choose what to test\n1 for Permulation\n2 for CheckOddInt\n3 for ValidateSmiley: ")
	switch input {
	case "1":
		result := getPermutationPromt()
		fmt.Println(result)
	case "2":
		result := getOddIntPromt()
		fmt.Println(result)

	case "3":
		result := getSmileyFacesPromt()
		fmt.Println(result)

	default:
		fmt.Println("That was not a valid option ...., please choose again")
		displayChoices()
	}
}
