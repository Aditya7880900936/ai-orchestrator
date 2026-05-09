package llm

import (
	"context"
	"os"

	openai "github.com/openai/openai-go"
)

var Client *openai.Client

func InitOpenAI() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(
		openai.WithAPIKey(apiKey),
	)

	Client = &client
}

func Generate(prompt string) (string, error) {
	resp, err := Client.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Model: "gpt-4.1-mini",
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}