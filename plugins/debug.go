package plugins

import (
	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/log"
)

func OnReady(_ *discordgo.Session, m *discordgo.Ready) {
	log.Log.Infof("Ready on Shard #%d", m.Shard)
}
