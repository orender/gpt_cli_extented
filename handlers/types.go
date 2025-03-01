package handlers

import openai "github.com/sashabaranov/go-openai"

// Message struct to represent a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// InitializeMessages returns initial system prompt message
func InitializeMessages() []openai.ChatCompletionMessage {
	systemPrompt := `You are a helpful assistant with an integrated multiplication function.
Your primary role is to engage in natural language conversations and answer user questions.`

	return []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
	}
}
