package main

import (
	"io"
	"log"
	"os"
)

func main() {
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading the input data by data-cleaning container %s: %v", inputData, err)
	}

	toBeSentData := "Output via Data Cleaning Container"

	io.WriteString(os.Stdout, toBeSentData)
}
