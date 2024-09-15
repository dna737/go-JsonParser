package main

import (
	"log"
	"testing"
)

func TestValidJson(t *testing.T){
	input :=  []string {`{"key": "value"}`, `{"key1": true, "key2": false, "key3": {}}`, `{"name": [1, 2, 3]}`}	

	for _, t := range input {
		if !validateJson(t) {
			log.Fatalf("Expected valid JSON for %v got invalid instead.", t)
		}
	}
}

func TestInvalidJson(t *testing.T) {
	input :=  []string {`
	{
	"key": "value"
	"key2": "value"
	}
	`, 
	`{
		"key1": true,
		"key2": false,
		true: null,
		null: "value",
		"key5": 101
	}`}	
	
	for _, t := range input {
		if validateJson(t) {
			log.Fatalf("Expected invalid JSON for %v got valid instead.", t)
		}
	}
}
