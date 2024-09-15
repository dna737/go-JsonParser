package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func seekInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
		os.Exit(1)
	}

	return text
}

func checkEnds(input string, stringRange int) bool {

	// TODO: only check the last ones as the full iteration is not required.
	for i := 0; i < stringRange ; i++ {
		char := string(input[i])
		if i == 0 && char != "{" {
			fmt.Println("Expected char to be '{', but instead got", char)
			return false
		}

		if i == stringRange - 1  && char != "}" {
			fmt.Println("Expected char to be '}', but instead got", char)
			return false
		}
	}

	return true
}

func isEmptyJson(input string, openIndex, closeIndex int) bool {
	return  strings.TrimSpace(input[openIndex + 1 : closeIndex]) == ""
}

func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}

func isValidEntity(entity string, isKey bool) bool {

	firstChar, lastChar := string(entity[0]), string(entity[len(entity) - 1]) 
	
	if isKey {
		return (firstChar == "\"" && lastChar == "\"")
	}

	return (
		firstChar == "\"" && lastChar == "\"" ||
		firstChar == "[" && lastChar == "]" ||
		entity == "true" || 
		entity == "false" ||
		entity == "null" || 
		isInt(entity) ||
		validateJson(entity))
}

func validateJson(text string) bool {

	input := strings.TrimSpace(text)

	stringRange := len(input)	

	if string(input[len(input) - 1]) == "\n" {
		stringRange -= 1
	}
	if !checkEnds(input, stringRange) {
		return false
	}

	openIndex, closeIndex := strings.Index(input, "{"), strings.LastIndex(input, "}")
	if isEmptyJson(input, openIndex, closeIndex) {
			return true
	}
	if strings.Count(input, ":") != strings.Count(input, ",") + 1 {
		return false
	}

	inputExcludingBraces := input[openIndex + 1 : closeIndex]

	//Extracting out the key and values for each pair:
	for _, pair := range strings.Split(inputExcludingBraces, ",") {
		if !strings.Contains(pair, ":") {
			return false
		}

		//Extracts kv from a pair and trims out the whitespace.
		key, value := func() (string, string) { parts := strings.Split(pair, ":"); return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]) }()

		if !isValidEntity(key, true) || !isValidEntity(value, false) {
			return false
		}
	}
	
	return true
}

func main() {
	input := seekInput()
	if validateJson(input) {
		fmt.Println("The JSON is valid.")
	} else {
		fmt.Println("The JSON is invalid.")
	}
}
