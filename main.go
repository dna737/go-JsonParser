package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func isValidEntity(entity string) bool {
	// fmt.Println(len(entity) - 1)
	return string(entity[0]) == "\"" //&& string(entity[len(entity) - 1 ]) == "\""
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

	//Extracting out the key and values:
	openIndex, closeIndex, colonIndex := strings.Index(input, "{"), strings.Index(input, "}"), strings.Index(input, ":")

	if isEmptyJson(input, openIndex, closeIndex) {
		return true
	}

	fmt.Println(closeIndex, openIndex)
	key, value := input[openIndex + 1 : colonIndex], input[colonIndex + 1 : closeIndex]
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	fmt.Println("key", key, "value", value)
	if !isValidEntity(key) || !isValidEntity(value) {
		fmt.Println("fail")
		return false
	}
	
	return true
}

func main() {
	input := seekInput()
	if validateJson(input) {
		fmt.Println("The JSON is valid.")
	}
}
