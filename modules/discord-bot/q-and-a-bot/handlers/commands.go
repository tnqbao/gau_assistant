package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tnqbao/gau_assistant/config/gemini_api"
	"log"
)

func HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {

	aiClient := gemini_api.NewAIClient()

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		log.Printf("Error acknowledging interaction: %v", err)
		return
	}

	if len(i.ApplicationCommandData().Options) == 0 {
		log.Printf("No options provided in the command.")
		return
	}

	input := i.ApplicationCommandData().Options[0].StringValue()
	response, err := aiClient.GetResponse(input)
	if err != nil {
		log.Printf("Error calling AI: %v", err)
		response = "Sorry, I couldn't process your request. Please try again later."
	} else {
		log.Printf("Raw AI response: %v", response)
	}

	content := fmt.Sprintf(response)
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	})
	if err != nil {
		log.Printf("Error editing response: %v", err)
	}
}
