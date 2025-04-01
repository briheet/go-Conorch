package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
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

	// Read the Prompt file contents
	prompt, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading the Prompt file: %v", err)
	}

	// Read the Data file that needs to be worked on
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatalf("Error reading the file data: %v", err)
	}

	promptData := string(prompt) + string(data)

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
				Content: string(promptData),
			}, {
				Role:    "user",
				Content: string(containersList),
			},
		},
	}

	// Marshal this payload
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new request
	req, err := http.NewRequest(http.MethodPost, "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}

	// Add required headers, Get the API Key via environment variable
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))

	// A new http client
	client := http.Client{}

	// Make a request, get the response from LLM back
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making API request: %v", err)
	}

	// Very Important, please defer to close
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var response LLMResponse

	// Unmarshal the response data
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Get all the containers required
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

	// Execute the containers which are suggested by the LLM, yeee-Haaa
	err = executeContainers(containerResponse)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func executeContainers(containerList []string) error {

	// Read the data file first, very important to get the non polluted data first
	// Build all the containers, run them, get their logs, move to the new container

	// Build the containers before executing with the help of a goroutine
	for _, container := range containerList {
		go func(container string) {
			buildCmd := exec.Command("docker", "build", "-t", container, "-f", fmt.Sprintf("tools/%s/%s.multistage", container, container), fmt.Sprintf("tools/%s/", container))

			fmt.Println("Building:", fmt.Sprintf("tools/%s/%s.multistage", container, container))
			if err := buildCmd.Run(); err != nil {
				log.Fatalf("Error building container %s: %v", container, err)
			}

		}(container)
	}

	for _, containerName := range containerList {

		fileData, err := os.ReadFile("data.txt")
		if err != nil {
			return fmt.Errorf("error reading data.txt: %v", err)
		}

		// Append asci for newline
		fileData = append(fileData, 0x04)

		// Remove the container if it exists
		rmCmd := exec.Command("docker", "rm", "-f", containerName)
		if err := rmCmd.Run(); err != nil {
			return fmt.Errorf("Error removing the already present container %s: %v", containerName, err)
		}

		// Run command
		runCmd := exec.Command("docker", "run", "--name", containerName, "-i", containerName)

		// Get Std input via StdingPipe
		stdin, err := runCmd.StdinPipe()
		if err != nil {
			return fmt.Errorf("error getting Stdin Pipe for %s: %w", containerName, err)
		}

		// Write the fcking data man
		go func(fileData []byte) {
			defer stdin.Close()
			io.WriteString(stdin, string(fileData))

		}(fileData)

		// Fking stdout
		stdout, err := runCmd.StdoutPipe()
		if err != nil {
			return fmt.Errorf("error inting the stdout in %s: %v", containerName, err)
		}

		// Start the command
		if err := runCmd.Start(); err != nil {
			return fmt.Errorf("error staring the runCmd in %s: %v", containerName, err)
		}

		// Delete previous contents of the file for storing new data
		if err := os.Truncate("data.txt", 0); err != nil {
			return fmt.Errorf("failed to trunicate file in %s: %v", containerName, err)
		}

		// Gather output
		output, err := io.ReadAll(stdout)
		if err != nil {
			return fmt.Errorf("error reading the output of the container in %s: %v", containerName, err)
		}

		// Open file with permissions
		resultFile, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("error opening the resultFile in %s: %v", containerName, err)
		}

		// Write the output
		_, err = resultFile.Write(output)
		if err != nil {
			return fmt.Errorf("error writing all the output in %s: %v", containerName, err)
		}

		if err := runCmd.Wait(); err != nil {
			return fmt.Errorf("error waiting for the runCmd to finish in %s: %v", containerName, err)

		}

	}

	return nil
}

func prettyPrint(body []byte) {

	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, body, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

}
