package api

import openai "github.com/sashabaranov/go-openai"

func CreateFunctions() []openai.FunctionDefinition {
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
