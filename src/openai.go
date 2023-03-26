package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func handleOpenAi(prompt string) (string, error) {
	// .env ファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("Error loading .env file: %w", err)
	}

	token := os.Getenv("OPENAI_TOKEN")

	c := openai.NewClient(token)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 200,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return "", fmt.Errorf("ChatCompletionStream error: %w", err)
	}
	defer stream.Close()

	var generatedText string
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return generatedText, nil
		}
		if err != nil {
			return "", fmt.Errorf("Stream error: %w", err)
		}
		generatedText += response.Choices[0].Delta.Content
	}
}
