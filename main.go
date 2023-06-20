package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sashabaranov/go-openai"
)

func main() {
	// Set up Telegram bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Set up OpenAI client
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// Set up updates channel
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Process incoming messages
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Get user input
		userInput := update.Message.Text

		// Generate response using GPT-4
		response, err := client.CreateCompletion(&openai.CompletionRequest{
			Model:     "gpt-4.0-turbo",
			Prompt:    userInput,
			MaxTokens: 50,
		})
		if err != nil {
			log.Fatal(err)
		}

		// Send response to user
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.Choices[0].Text)
		_, err = bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
