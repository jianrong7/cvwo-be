package initializers

import (
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func OpenAiClient() *gogpt.Client {
	gpt := gogpt.NewClient(os.Getenv("OPENAI_API"))
	return gpt
}