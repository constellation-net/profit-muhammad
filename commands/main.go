package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/log"
)

var (
	// Commands is used to store app command info before it is then registered via Register.
	Commands []*discordgo.ApplicationCommand
	// Handlers is akin to Commands, except maps each command's name to a handler that executes the command.
	Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
	// Registered contains every command that has been registered during the current session, this is used to deregister commands when Muhammad shuts down.
	Registered []*discordgo.ApplicationCommand
)

// DispatchHandler is the middleware that sends incoming commands to their respective handler.
// If no handler is registered, a warning message is given that the handler is missing.
func DispatchHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	if h, ok := Handlers[e.ApplicationCommandData().Name]; ok {
		h(s, e)
	} else {
		log.Log.Warnf("Received command '%s', but no handler is present.", e.ApplicationCommandData().Name)
	}
}

// Register tells Discord about all the interactions in Commands, and adds them to Registered so we can keep track of them.
// Plugins are expected to add their commands' info to Commands BEFORE Register is called by using an init() function in their file.
func Register(s *discordgo.Session) {
	for _, c := range Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", c)
		log.Error(err, fmt.Sprintf("COMMAND_CREATE (%s)", c.Name))

		Registered = append(Registered, cmd)
	}

	log.Log.Debugf("Registered %d commands", len(Commands))
}

// Deregister does the opposite of Register. This ensures that Muhammad's commands disappear whenever it shuts down.
// This helps deal with any rogue commands that may be removed from the bot after being registered.
func Deregister(s *discordgo.Session) {
	for _, c := range Registered {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", c.ID)
		log.Error(err, fmt.Sprintf("COMMAND_REMOVE (%s)", c.Name))
	}
}
