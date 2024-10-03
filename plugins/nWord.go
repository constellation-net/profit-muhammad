package plugins

import (
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/constellation-net/profit-muhammad/data"
)

var (
	Cooldown  = 60 * time.Second
	UserLocks = map[string]*sync.Mutex{}
)

func NWordCounter(s *discordgo.Session, e *discordgo.MessageCreate) {
	lock, ok := UserLocks[e.Author.ID]
	if !ok {
		lock = &sync.Mutex{}
		UserLocks[e.Author.ID] = lock
	}
	canLock := lock.TryLock()
	if !canLock {
		return
	}

	data.IncrementUserScore("nword", e.Author.ID)

	time.Sleep(Cooldown)
	lock.Unlock()
}
