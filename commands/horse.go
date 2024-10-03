package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	// Emojis is a list of emojis used when fish reacting to a message.
	HorseEmojis = []string{
		"🐎",
		"🐴",
		"🏇",
		"🎠",
		"🦄",
	}
)

func init() {
	name := "Horse react"

	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name: name,
		Type: discordgo.MessageApplicationCommand,
	})
	Handlers[name] = HorseReact
}

// HorseReact is a message command that reacts to a message with a bunch of horse emojis.
func HorseReact(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Horse react this man!",
		},
	})

	for _, m := range e.ApplicationCommandData().Resolved.Messages {
		for _, emj := range HorseEmojis {
			s.MessageReactionAdd(m.ChannelID, m.ID, emj)
		}
	}
}
