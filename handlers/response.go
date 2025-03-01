package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
)

const requestTimeout = 60 * time.Second

var functionHandler = &FunctionHandler{}

func StreamResponse(messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition, apiKey string) {
	client := openai.NewClient(apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4oMini,
		Messages:  messages,
		Stream:    true,
		Functions: functions,
	}

	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		log.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				return
			}
			log.Printf("Stream Recv error: %v\n", err)
			return
		}

		HandleChunk(response)
	}
}

func HandleChunk(response openai.ChatCompletionStreamResponse) {
	for _, choice := range response.Choices {
		if choice.Delta.Content != "" {
			fmt.Print(color.WhiteString(choice.Delta.Content))
		}

		if choice.Delta.FunctionCall != nil {
			functionHandler.HandleFunctionCall(choice.Delta.FunctionCall)
		}
	}
}
