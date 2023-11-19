package main

func validateEmptyStr(str string) bool {
	return str != ""
}

func validateSpaceInString(str string) bool {
	result := true
	for _, character := range str {
		if character == ' ' {
			result = false
		}
	}
	return result
}
