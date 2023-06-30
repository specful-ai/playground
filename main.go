package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/openai/openai-go/v1"
)

func main() {
	// Initialize OpenAI API client
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// Initialize Telegram bot API
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	// Set bot to receive updates from all messages
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Process incoming messages
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Process user input in natural language
		if update.Message.From.UserName != bot.Self.UserName {
			response, err := client.Completions(
				openai.CompletionRequest{
					Model: "text-davinci-002",
					Prompt: update.Message.Text,
					MaxTokens: 60,
					Temperature: 0.5,
				},
			)
			if err != nil {
				log.Panic(err)
			}

			// Send response to user
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.Choices[0].Text)
			bot.Send(msg)
		}

		// Process formatted commands from agent
		if update.Message.From.UserName == bot.Self.UserName {
			switch update.Message.Command() {
			case "search":
				// Search on the internet
				// TODO: Implement search functionality
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Searching...")
				bot.Send(msg)
			case "action":
				// Take some action
				// TODO: Implement action functionality
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Taking action...")
				bot.Send(msg)
			}
		} else {
			// Process formatted commands from assistant
			switch update.Message.Command() {
			case "agent":
				// Call agent to get more information
				// TODO: Implement agent functionality
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Calling agent...")
				bot.Send(msg)
			}
		}
	}
}
