package plugins

import (
	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/plugins/crazy"
	"github.com/Jack-Gledhill/profit-muhammad/plugins/entrance"
	"github.com/Jack-Gledhill/profit-muhammad/plugins/mention"
	"github.com/Jack-Gledhill/profit-muhammad/plugins/real"
)

func Register(s *discordgo.Session) {
	s.AddHandler(crazy.Crazy)
	s.AddHandler(entrance.Entrance)
	s.AddHandler(mention.Mention)
	s.AddHandler(real.Real)
	s.AddHandler(OnReady)
}
