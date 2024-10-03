package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/constellation-net/profit-muhammad/commands"
	"github.com/constellation-net/profit-muhammad/config"
	"github.com/constellation-net/profit-muhammad/data"
	"github.com/constellation-net/profit-muhammad/events"
	"github.com/constellation-net/profit-muhammad/log"
)

func main() {
	defer data.Disconnect()

	client, err := discordgo.New("Bot " + config.Config.Discord.Token)
	log.Error(err, "CLIENT_CREATE", true)

	client.AddHandler(commands.DispatchHandler)
	log.Log.Debug("Command dispatch handler created")

	client.Identify.Intents = discordgo.IntentsGuildMessages + discordgo.IntentsGuilds + discordgo.IntentsGuildMembers
	client.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name:  config.Config.Discord.Activity.Name,
			State: config.Config.Discord.Activity.State,
			Type:  discordgo.ActivityType(config.Config.Discord.Activity.Type),
		},
		Status: config.Config.Discord.Activity.Status,
		AFK:    false,
	}

	err = client.Open()
	log.Error(err, "CLIENT_START", true)
	defer client.Close()
	log.Log.Debug("Client connection started")

	commands.Register(client)
	defer commands.Deregister(client)

	events.Register(client)

	log.Log.Info("Bot is now running, use CTRL-C to stop...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
