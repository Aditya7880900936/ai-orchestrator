package llm

import (
	"context"
	"os"

	"google.golang.org/genai"
)

type LLMClient interface {
	Generate(prompt string) (string, error)
}

type GeminiClient struct {
	client *genai.Client
}

var Client LLMClient

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

	Client = &GeminiClient{
		client: client,
	}

	return nil
}

func (g *GeminiClient) Generate(prompt string) (string, error) {

	ctx := context.Background()

	resp, err := g.client.Models.GenerateContent(
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

func Generate(prompt string) (string, error) {
	return Client.Generate(prompt)
}
