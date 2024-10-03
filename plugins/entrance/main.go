package entrance

import (
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/utils"
)

var (
	Message = "Hello thank you for calling Microsoft tech support, my name is Mohammad and I will help fix computer"
)

func Entrance(s *discordgo.Session, e *discordgo.GuildCreate) {
	// Check join time, if later than a minute ago, it's likely just the API giving us guild info
	now := time.Now()
	if now.Sub(e.JoinedAt) > 60*time.Second {
		return
	}

	if len(e.Channels) > 0 {
		cID := e.Channels[0].ID

		utils.SimulateTyping(s, cID, Message)
		s.ChannelMessageSend(cID, Message)
	}
}
