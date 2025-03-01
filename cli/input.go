package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gpt-cli/handlers"

	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
)

func RunCLI(messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition, apiKey string) {
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
		handlers.StreamResponse(messages, functions, apiKey)
	}
}
