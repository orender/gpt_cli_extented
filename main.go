package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

const (
	modelName    = openai.GPT4oMini
	systemPrompt = `You are a helpful assistant with an integrated multiplication function.
Your primary role is to engage in natural language conversations and answer user questions.

You have access to a function called 'multiply' that performs multiplication.
You should ONLY use the 'multiply' function when the user's input explicitly and directly requests a multiplication calculation, such as:
- "What is 5 times 7?"
- "Calculate 12 * 3."
- "Multiply 8 by 9."
- "What is the result of 10 * 10?"
- "9*9"

For all other conversations and questions, respond using natural language.
Do not use the 'multiply' function unless the user's input clearly indicates a multiplication operation.

Maintain context throughout the conversation and only use the function when it is absolutely necessary and directly requested.
If there is any doubt, respond with normal language.`
	requestTimeout = 60 * time.Second
)

var functionArgsBuffer strings.Builder
var functionName string

func main() {
	apiKey, err := loadAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleSystem, Content: systemPrompt}}
	functions := createFunctions()

	runCLI(messages, functions, apiKey)
}

func loadAPIKey() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY not set in .env")
	}
	return apiKey, nil
}

func createFunctions() []openai.FunctionDefinition {
	return []openai.FunctionDefinition{
		{
			Name: "multiply",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"a": map[string]interface{}{"type": "integer", "description": "The first number to multiply."},
					"b": map[string]interface{}{"type": "integer", "description": "The second number to multiply."},
				},
				"required": []string{"a", "b"},
			},
		},
	}
}

func runCLI(messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition, apiKey string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(color.CyanString("GPT-4o-mini CLI - Type your message or 'exit' to quit:"))

	for {
		color.Green("\nYou: ")
		if !scanner.Scan() {
			break
		}
		userInput := scanner.Text()

		if strings.EqualFold(userInput, "exit") {
			fmt.Println(color.CyanString("Goodbye!"))
			break
		}

		fmt.Println(color.YellowString("\nGPT-4o-mini:\n"))

		messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: userInput})
		streamResponse(messages, functions, apiKey)
	}
}

func streamResponse(messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition, apiKey string) {
	client := openai.NewClient(apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := openai.ChatCompletionRequest{
		Model:     modelName,
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

		handleChunk(response)
	}
}

func handleChunk(response openai.ChatCompletionStreamResponse) {
	for _, choice := range response.Choices {
		if choice.Delta.Content != "" {
			fmt.Print(color.WhiteString(choice.Delta.Content))
		}
		if choice.Delta.FunctionCall != nil {
			handleFunctionCall(choice.Delta.FunctionCall)
		}
	}
}

func handleFunctionCall(funcCall *openai.FunctionCall) {
	if funcCall.Name != "" {
		functionName = funcCall.Name
	}

	if funcCall.Arguments != "" {
		functionArgsBuffer.WriteString(funcCall.Arguments)
	}

	if strings.HasSuffix(funcCall.Arguments, "}") {
		args := map[string]int{}
		if err := json.Unmarshal([]byte(functionArgsBuffer.String()), &args); err != nil {
			log.Printf("Error unmarshaling function arguments: %v", err)
			functionArgsBuffer.Reset()
			return
		}

		if functionName != "" {
			fmt.Printf("\n%s: %s\n", color.CyanString("Function Name"), functionName)

			switch functionName {
			case "multiply":
				result := args["a"] * args["b"]
				fmt.Printf("\n%s: %d\n", color.MagentaString("Result"), result)
			default:
				fmt.Printf("\n%s\n", color.RedString("Unrecognized func name recieved"))
			}

			functionName = ""
		}

		functionArgsBuffer.Reset()
	}
}
