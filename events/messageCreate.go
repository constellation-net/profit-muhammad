package events

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/constellation-net/profit-muhammad/config"
	"github.com/constellation-net/profit-muhammad/plugins"
	"github.com/constellation-net/profit-muhammad/utils"
)

var (
	Triggers []*TriggerConfig = []*TriggerConfig{
		{
			Cooldown: 300 * time.Second,
			Responses: []string{
				"Crazy?",
				"I was crazy once",
				"They locked me in a room",
				"A rubber room",
				"A rubber room with rats",
				"Rats make me crazy",
			},
			Triggers: []string{
				"crazy",
			},
		},
		{
			Cooldown: 60 * time.Second,
			Responses: []string{
				"Whaddaya want?",
				"The fuck you want little shit?",
				"WHAT!?",
				"Fuck off",
				"The fuck you want?",
				"Please shut up",
			},
			Triggers: []string{
				fmt.Sprintf("<@%s>", config.Config.Discord.BotID),
			},
		},
		{
			Callback: plugins.NWordCounter,
			Triggers: []string{},
		},
	}
)

type TriggerConfig struct {
	Callback  func(*discordgo.Session, *discordgo.MessageCreate)
	Cooldown  time.Duration
	Lock      sync.Mutex
	Responses []string
	Triggers  []string
}

func MessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	// Ignore bots
	if e.Author.Bot {
		return
	}

	for _, w := range strings.Split(strings.ToLower(e.Content), " ") {
		for _, t := range Triggers {
			if utils.SliceContains(t.Triggers, w) {
				canLock := t.Lock.TryLock()
				// If we can't lock, we want to abort rather than wait to continue
				if !canLock {
					return
				}

				// If length of Responses is 0 then this loop won't run
				for i := 0; i < len(t.Responses); i++ {
					m := t.Responses[i]
					utils.SimulateTyping(s, e.ChannelID, m)
					s.ChannelMessageSend(e.ChannelID, m)
				}

				if t.Callback != nil {
					// Runs the callback function concurrently, so it doesn't block this loop
					// This is good for callbacks that take a long time or have their own cooldown system in place
					go t.Callback(s, e)
				}

				// Sleeping here effectively creates a cooldown, since the lock remains on for longer
				// If t.Cooldown is not specified, it should default to 0, which is effectively no cooldown
				time.Sleep(t.Cooldown)
				t.Lock.Unlock()

				// Stops the same message triggering multiple autoreploes
				break
			}
		}
	}
}
