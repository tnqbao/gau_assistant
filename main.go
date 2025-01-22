package main

import (
	"github.com/joho/godotenv"
	"github.com/tnqbao/gau_assistant/modules/discord-bot"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("No Discord token provided")
	}

	bot, err := discord_bot.NewDiscordBot(token)
	if err != nil {
		log.Fatalf("Error creating Discord bot: %v", err)
	}

	err = bot.Run()
	if err != nil {
		log.Fatalf("Error running Discord bot: %v", err)
	}

	select {}
}
