package events

import (
	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session) {
	s.AddHandler(GuildCreate)
	s.AddHandler(MessageCreate)
	s.AddHandler(Ready)
}
