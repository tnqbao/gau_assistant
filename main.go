package main

import (
	finannce_management "github.com/tnqbao/gau_assistant/modules/telegram-bot/finannce-management"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/tnqbao/gau_assistant/modules/discord-bot/q-and-a-bot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	finannce_management.InitDB()

	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("No Discord token provided")
	}

	discordBot, err := q_and_a_bot.NewDiscordBot(discordToken)
	if err != nil {
		log.Fatalf("Error creating Discord bot: %v", err)
	}

	go func() {
		err := discordBot.Run()
		if err != nil {
			log.Fatalf("Error running Discord bot: %v", err)
		}
	}()

	telegramBot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal("Lỗi khởi tạo bot Telegram:", err)
	}
	telegramBot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := telegramBot.GetUpdatesChan(u)

	for update := range updates {
		go finannce_management.HandleMessage(telegramBot, update)
	}

	select {}
}
