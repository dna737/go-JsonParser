package main

import (
	"log"
	"testing"
)

func TestValidJson(t *testing.T){
	input :=  []string {`{"name": "Adnan"}`, `{}   `, `{   }`}	

	for _, t := range input {
		if !validateJson(t) {
			log.Fatalf("Expected valid JSON for %v got invalid instead.", t)
		}
	}
}

func TestInvalidJson(t *testing.T) {
	input :=  []string {`{`, `}   `, `{{`, `}}}`}	

	
	for _, t := range input {
		if validateJson(t) {
			log.Fatalf("Expected invalid JSON for %v got valid instead.", t)
		}
	}
}
