package main

import (
	"log"
	"os"

	"github.com/briheet/ai-orchestrator/cli"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewCli()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
