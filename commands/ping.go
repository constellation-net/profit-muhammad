package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	name := "ping"

	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        name,
		Description: "Pong! 🏓",
	})
	Handlers[name] = Ping
}

// Ping can be used to test Muhammad's responsiveness, and also provides the most recent heartbeat latency.
func Ping(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Pong! 🏓\n☁️ **Heartbeat latency:** %d ms", s.HeartbeatLatency().Milliseconds()),
		},
	})
}
