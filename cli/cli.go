package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func NewCli() *cli.App {
	return &cli.App{

		Name:     "AI Orchestrator",
		HelpName: "ai-orc",
		Usage:    "Container orchestrator coupled with AI recommendation",
		Action: func(*cli.Context) error {
			fmt.Println("Please specify the command to be executed")
			return nil
		},

		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Briheet Singh Yadav",
				Email: "briheetyadav@gmail.com",
			},
		},

		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Get container recommendation",
				Action: orchestrateCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "Path to the text file to analyze",
						Required: true,
					},
				},
			},
		},
		Suggest:              true,
		EnableBashCompletion: true,
	}
}

type Payload struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
	XGroq             XGroq    `json:"x_groq"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     any     `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	QueueTime        float64 `json:"queue_time"`
	PromptTokens     int     `json:"prompt_tokens"`
	PromptTime       float64 `json:"prompt_time"`
	CompletionTokens int     `json:"completion_tokens"`
	CompletionTime   float64 `json:"completion_time"`
	TotalTokens      int     `json:"total_tokens"`
	TotalTime        float64 `json:"total_time"`
}

type XGroq struct {
	ID string `json:"id"`
}

func orchestrateCommand(c *cli.Context) error {

	// ai-orc orch -f now.txt
	filePath := c.String("file")

	// Read file contents
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Get all containers
	containersList, err := os.ReadFile("containersList.txt")
	if err != nil {
		log.Fatal("Error reading the containers list:", err)
	}

	payload := Payload{
		Model: "llama-3.3-70b-versatile", // Hardcoding this for now
		Messages: []Message{
			{
				Role:    "user",
				Content: "You are an AI assistant that helps select predefined containers for user tasks.",
			}, {
				Role:    "user",
				Content: "When a user provides a request, respond with container names in increseasing order if more than one containers are required",
			}, {
				Role:    "user",
				Content: "Here is the task and the names of the pre defined containers we have, only return container names with new lines",
			}, {
				Role:    "user",
				Content: string(fileData),
			}, {
				Role:    "user",
				Content: string(containersList),
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making API request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	prettyPrint(body)

	var response LLMResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	var containerResponse []string
	if len(response.Choices) > 0 {

		for i := 0; i < len(response.Choices); i++ {
			containerresponse := response.Choices[i].Message.Content

			lines := strings.Split(containerresponse, "\n")

			for _, line := range lines {
				strings.TrimSpace(line)
				containerResponse = append(containerResponse, line)
			}

		}

	} else {
		log.Fatal("No choices recommended by the model")
	}

	err = executeContainers(containerResponse)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func executeContainers(containerList []string) error {
	return nil
}

func prettyPrint(body []byte) {

	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, body, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

}
