package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	// Emojis is a list of emojis used when fish reacting to a message.
	Emojis = []string{
		"🐟",
		"🐠",
		"🐡",
		"🦐",
		"🦈",
		"🎣",
	}
)

func init() {
	name := "Fish react"

	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name: name,
		Type: discordgo.MessageApplicationCommand,
	})
	Handlers[name] = FishReact
}

// FishReact is a message command that reacts to a message with a bunch of fish emojis.
func FishReact(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Fish react this man!",
		},
	})

	for _, m := range e.ApplicationCommandData().Resolved.Messages {
		for _, emj := range Emojis {
			s.MessageReactionAdd(m.ChannelID, m.ID, emj)
		}
	}
}
