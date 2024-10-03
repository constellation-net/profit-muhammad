package mention

import (
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/utils"
)

var (
	Lock     = sync.Mutex{}
	Messages = []string{
		"Whaddaya want?",
		"The fuck you want little shit?",
		"WHAT!?",
		"Fuck off",
		"The fuck do you want?",
	}
	Triggers = []string{
		"<@1176638378570174557>",
	}
)

func Mention(s *discordgo.Session, e *discordgo.MessageCreate) {
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

			m := utils.RandSliceItem(Messages)
			utils.SimulateTyping(s, e.ChannelID, m)
			s.ChannelMessageSend(e.ChannelID, m)

			// Stops it running twice if multiple triggers are present
			Lock.Unlock()
			break
		}
	}
}
