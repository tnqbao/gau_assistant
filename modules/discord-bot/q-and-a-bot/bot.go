package q_and_a_bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tnqbao/gau_assistant/modules/discord-bot/q-and-a-bot/handlers"
	"log"
)

type DiscordBot struct {
	Session *discordgo.Session
}

func NewDiscordBot(token string) (*DiscordBot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &DiscordBot{Session: session}

	session.AddHandler(bot.ready)
	session.AddHandler(handlers.HandleCommand)

	return bot, nil
}

func (bot *DiscordBot) ready(s *discordgo.Session, r *discordgo.Ready) {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "ask",
		Description: "Ask the AI anything",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "question",
				Description: "Your question",
				Required:    true,
			},
		},
	})
	if err != nil {
		log.Fatalf("Error creating command: %v", err)
	}
	log.Println("Slash command '/ask' created successfully.")
}

func (bot *DiscordBot) Run() error {
	err := bot.Session.Open()
	if err != nil {
		return err
	}
	log.Println("Discord bot is running.")
	return nil
}

func (bot *DiscordBot) Close() {
	bot.Session.Close()
}
