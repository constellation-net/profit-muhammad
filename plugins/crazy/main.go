package crazy

import (
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/utils"
)

var (
	Cooldown = 300 * time.Second
	Lock     = sync.Mutex{}
	Messages = []string{
		"Crazy?",
		"I was crazy once",
		"They locked me in a room",
		"A rubber room",
		"A rubber room with rats",
		"Rats make me crazy",
	}
	Triggers = []string{
		"crazy",
	}
)

func Crazy(s *discordgo.Session, e *discordgo.MessageCreate) {
	// Ignore bots
	if e.Author.Bot {
		return
	}

	for _, w := range strings.Split(strings.ToLower(e.Content), " ") {
		if utils.SliceContains(Triggers, w) {
			canLock := Lock.TryLock()
			// If we can't lock, we want to abort rather than wait to continue
			if !canLock {
				return
			}

			for i := 0; i < len(Messages); i++ {
				m := Messages[i]
				utils.SimulateTyping(s, e.ChannelID, m)
				s.ChannelMessageSend(e.ChannelID, m)
			}

			// Sleeping here effectively creates a cooldown, since the lock remains on for even longer
			time.Sleep(Cooldown)
			Lock.Unlock()

			// Stops it running twice if multiple triggers are present
			break
		}
	}
}
