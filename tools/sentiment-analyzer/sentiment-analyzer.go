package main

import (
	"io"
	"log"
	"os"
)

func main() {

	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading the input data by container %s: %v", inputData, err)
	}

	toBeSentData := "Output via Sentiment Analyzer Container"

	io.WriteString(os.Stdout, toBeSentData)

}
