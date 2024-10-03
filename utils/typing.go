package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	// CPW is the assumed average characters per word. Combines with WPM to form a typing speed in characters per minute (CPM).
	CPW float64 = 5
	// WPM - Mohammad's simulated typing speed (in words per minute), affects how long Mohammad is typing before sending a message.
	WPM float64 = 70
	// MaxTypingTime is the maximum time Mohammad will spend typing before sending a message, regardless of how long it is.
	MaxTypingTime float64 = 10 // Discord's maximum typing time is 10 seconds
	// NoiseLimit is the limit of additional noise that can be added to the calculated typing time.
	NoiseLimit float64 = 1
)

var (
	// CPS is the calculated typing speed in characters per second
	CPS float64 = WPM * CPW / 60
)

// CalcTypeTime will calculate how long Muhammad should spend typing before sending a message, based on how many characters are in the message.
// For additional realism, a random float value of noise is added. The maximum value that this noise can be is NoiseLimit.
func CalcTypeTime(m string) float64 {
	// Calculate typing time based on typing speed
	charCount := float64(len(strings.Split(m, "")))
	typeTime := charCount / CPS

	if typeTime <= MaxTypingTime {
		// rand.Float64() returns a float between 0 and 1, hence to scale that up to NoiseLimit, we just need to multiply by NoiseLimit
		noise := rand.Float64() * NoiseLimit
		return typeTime + noise
	} else {
		return MaxTypingTime
	}
}

// SimulateTyping makes Muhammad type for a certain amount of time before sending a message.
// This makes him feel a bit more real, since he doesn't send messages instantly.
// All this function does is start typing in a channel, waits a few seconds, then returns so that the message can be sent.
// This function does NOT actually send the message, it's intended that the message is sent by the caller immediately after this function returns.
func SimulateTyping(s *discordgo.Session, cID string, m string) error {
	err := s.ChannelTyping(cID)
	if err != nil {
		return err
	}

	waitTime := CalcTypeTime(m)
	time.Sleep(time.Duration(waitTime) * time.Second)
	return nil
}
