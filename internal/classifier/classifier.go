package classifier

import (
	"context"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

type Classifier struct {
	client *openai.Client
}

func NewClassifier(apiKey string) *Classifier {
	client := openai.NewClient(apiKey)
	return &Classifier{client: client}
}

func (c *Classifier) ClassifyEmail(emailContent string) (string, error) {
	ctx := context.Background()
	response, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are an email classifier. Classify the following email content as 'primary', 'Social', 'Marketing' or 'spam'.",
			},
			{
				Role:    "user",
				Content: emailContent,
			},
		},
	})

	if err != nil {
		log.Printf("Error classifying email: %v", err)
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
