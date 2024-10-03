package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/utils"
)

var (
	// Names maps an author's ID to their preferred name.
	Names = map[string]string{
		"269758783557730314": "Jack Gledhill",
		"431130121198632963": "Aidan Bailey",
		"618879980167757834": "Eevee Cash",
		"383532679766736897": "Jacob Wilson",
		"816842006647144458": "Freddie Elson",
	}
	// QuotesChannelID is the ID of the channel that quotes are sent to.
	QuotesChannelID = "1187737567953702922"
)

func init() {
	slashName := "quote"
	msgName := "Add to quotes"

	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        slashName,
		Description: "Quote a message",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "author",
				Description: "Whoever wrote the quote.",
				Required:    true,
				Type:        3,
			},
			{
				Name:        "date",
				Description: "The year that the quote was said.",
				Required:    true,
				Type:        4,
			},
			{
				Name:        "quote",
				Description: "The quote, duh.",
				Required:    true,
				Type:        3,
			},
		},
	})
	Handlers[slashName] = SlashQuote

	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name: msgName,
		Type: discordgo.MessageApplicationCommand,
	})
	Handlers[msgName] = MessageQuote
}

// SlashQuote is a slash command for quoting something a person has said.
// The message will be sent to QuotesChannelID in the following format:
//
//	"{quote}" - {author}, {date}
//
// Options
// -------
// quote  (string)  : the sentence(s) to be quoted.
// author (string)  : the person who said the quote.
// date   (integer) : the year that the quote was made.
func SlashQuote(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "OK",
		},
	})

	// Convert command options to map
	options := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(e.ApplicationCommandData().Options))
	for _, opt := range e.ApplicationCommandData().Options {
		options[opt.Name] = opt
	}

	m := fmt.Sprintf("\"%s\" - %s, %d", options["quote"].Value, options["author"].Value, int(options["date"].Value.(float64)))
	utils.SimulateTyping(s, QuotesChannelID, m)
	s.ChannelMessageSend(QuotesChannelID, m)
}

// MessageQuote is the message command equivalent of SlashQuote. Rather than having command options, this gets all the necessary information from the selected message.
// It tries to use Names to identify the full name of the author, but will default to the author's username if they're not in the map.
func MessageQuote(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "OK",
		},
	})

	for _, m := range e.ApplicationCommandData().Resolved.Messages {
		// Try get full name, default to username
		name, ok := Names[m.Author.ID]
		if !ok {
			name = m.Author.Username
		}

		// Send message to quotes channel
		msg := fmt.Sprintf("\"%s\" - %s, %d", m.Content, name, m.Timestamp.Year())
		utils.SimulateTyping(s, QuotesChannelID, msg)
		s.ChannelMessageSend(QuotesChannelID, msg)
	}
}
