package real

import (
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/utils"
)

var (
	Cooldown        = 300 * time.Second
	Lock            = sync.Mutex{}
	MessageInterval = 1 * time.Second
	Stages          = [][]string{
		{
			"What was that Eevee?",
			"What did you say Eevee?",
		},
		{
			"R-re-real?",
			"Real you say?",
			"R...r...real?",
		},
		{
			"i^2",
			"**i^2**",
			"# i^2",
		},
	}
	Triggers = []string{
		"real",
	}
	Users = []string{
		"269758783557730314",
	}
)

func Real(s *discordgo.Session, e *discordgo.MessageCreate) {
	// This plugin works on a whitelist system
	if !utils.SliceContains(Users, e.Author.ID) {
		return
	}

	// Loop over every word in the message
	for _, w := range strings.Split(strings.ToLower(e.Content), " ") {
		if utils.SliceContains(Triggers, w) {
			canLock := Lock.TryLock()
			// If we can't lock, we want to abort rather than wait to continue
			if !canLock {
				return
			}

			for _, stage := range Stages {
				m := utils.RandSliceItem(stage)
				utils.SimulateTyping(s, e.ChannelID, m)
				s.ChannelMessageSend(e.ChannelID, m)

				time.Sleep(MessageInterval)
			}

			// Sleeping here effectively creates a cooldown, since the lock remains on for even longer
			time.Sleep(Cooldown)
			Lock.Unlock()

			// Stops it running twice if multiple triggers are present
			break
		}
	}
}
