package llm

import (
	"context"
	"os"

	"google.golang.org/genai"
)

var Client *genai.Client

func InitGemini() error {
	ctx := context.Background()

	client, err := genai.NewClient(
		ctx,
		&genai.ClientConfig{
			APIKey: os.Getenv("GEMINI_API_KEY"),
		},
	)

	if err != nil {
		return err
	}

	Client = client
	return nil
}

func Generate(prompt string) (string, error) {
	ctx := context.Background()

	resp, err := Client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
