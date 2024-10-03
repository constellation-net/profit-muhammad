package events

import (
	"github.com/bwmarrin/discordgo"

	"github.com/constellation-net/profit-muhammad/log"
)

func Ready(_ *discordgo.Session, m *discordgo.Ready) {
	log.Log.Infof("Ready on Shard #%d", m.Shard)
}
