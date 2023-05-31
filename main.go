package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/openai/openai-go/v1"
)

func main() {
	// Get OpenAI API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set OPENAI_API_KEY environment variable")
		return
	}

	// Create OpenAI client
	client, err := openai.NewClient(apiKey)
	if err != nil {
		fmt.Println("Error creating OpenAI client:", err)
		return
	}

	// Get user input
	var userInput string
	fmt.Print("Enter a natural language command: ")
	fmt.Scanln(&userInput)

	// Call OpenAI GPT-3 API to translate natural language to shell commands
	completion, err := client.Completions.Create(
		"text",
		&openai.CompletionRequest{
			Prompt:    userInput,
			Model:     "text-davinci-002",
			MaxTokens: 60,
		},
	)
	if err != nil {
		fmt.Println("Error calling OpenAI API:", err)
		return
	}

	// Extract shell commands from OpenAI API response
	shellCommands := strings.TrimSpace(completion.Choices[0].Text)

	// Execute shell commands
	cmd := exec.Command("sh", "-c", shellCommands)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing shell commands:", err)
		return
	}
}
