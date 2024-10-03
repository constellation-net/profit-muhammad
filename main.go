package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/Jack-Gledhill/profit-muhammad/commands"
	"github.com/Jack-Gledhill/profit-muhammad/log"
	"github.com/Jack-Gledhill/profit-muhammad/plugins"
)

func main() {
	client, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	log.Error(err, "CLIENT_CREATE", true)

	client.AddHandler(commands.DispatchHandler)
	log.Log.Debug("Command dispatch handler created")

	client.Identify.Intents = discordgo.IntentsGuildMessages + discordgo.IntentsGuilds + discordgo.IntentsGuildMembers
	client.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name:  "custom",
			State: "Slayyyy",
			Type:  4,
		},
		Status: "online",
		AFK:    false,
	}

	err = client.Open()
	log.Error(err, "CLIENT_START", true)
	defer client.Close()
	log.Log.Debug("Client connection started")

	commands.Register(client)
	defer commands.Deregister(client)

	plugins.Register(client)

	log.Log.Info("Bot is now running, use CTRL-C to stop...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
