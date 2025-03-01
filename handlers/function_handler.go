package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
)

type FunctionHandler struct {
	Name       string
	ArgsBuffer strings.Builder
}

func (fh *FunctionHandler) HandleFunctionCall(funcCall *openai.FunctionCall) {
	if funcCall.Name != "" {
		fh.Name = funcCall.Name
	}

	if funcCall.Arguments != "" {
		fh.ArgsBuffer.WriteString(funcCall.Arguments)
	}

	if strings.HasSuffix(funcCall.Arguments, "}") {
		args := map[string]int{}
		if err := json.Unmarshal([]byte(fh.ArgsBuffer.String()), &args); err != nil {
			log.Printf("Error unmarshaling function arguments: %v", err)
			fh.ArgsBuffer.Reset()
			return
		}

		if fh.Name != "" {
			fmt.Printf("\n%s: %s\n", color.CyanString("Function Name"), fh.Name)

			switch fh.Name {
			case "multiply":
				result := args["a"] * args["b"]
				fmt.Printf("\n%s: %d\n", color.MagentaString("Result"), result)
			default:
				fmt.Println(color.RedString("Unrecognized function name received"))
			}
			fh.Name = ""
		}

		fh.ArgsBuffer.Reset()
	}
}
