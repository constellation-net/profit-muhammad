package events

import (
	"github.com/bwmarrin/discordgo"

	"github.com/constellation-net/profit-muhammad/utils"
)

const (
	JoinMessage = "Hello thank you for calling Microsoft tech support, my name is Muhammad and I will help fix computer"
)

func GuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	// This will technically run whenever the bot starts up because of how the Discord API works
	// But in this case, I'd say that's a feature

	if len(e.Channels) > 0 {
		cID := e.Channels[0].ID

		utils.SimulateTyping(s, cID, JoinMessage)
		s.ChannelMessageSend(cID, JoinMessage)
	}
}
