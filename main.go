package main

import (
	"log"

	"gpt-cli/api"
	"gpt-cli/cli"
	"gpt-cli/config"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	apiKey, err := config.LoadAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: api.SystemPrompt},
	}

	functions := api.CreateFunctions()

	cli.RunCLI(messages, functions, apiKey)
}
