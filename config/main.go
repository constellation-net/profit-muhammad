package config

import (
	"os"
	"strconv"
)

var (
	Config = &ConfigType{
		Discord: DiscordConfig{
			Activity: DiscordActivityConfig{
				Name:  os.Getenv("DISCORD_ACTIVITY_NAME"),
				State: os.Getenv("DISCORD_ACTIVITY_STATE"),
				Type: func() int {
					if v, err := strconv.Atoi(os.Getenv("DISCORD_ACTIVITY_TYPE")); err == nil {
						return v
					}
					return 0
				}(),
				Status: os.Getenv("DISCORD_ACTIVITY_STATUS"),
			},
			BotID: os.Getenv("DISCORD_BOT_ID"),
			Token: os.Getenv("DISCORD_TOKEN"),
		},
		Environment: os.Getenv("ENVIRONMENT"),
		Mongo: MongoConfig{
			URL: os.Getenv("MONGO_URL"),
		},
	}
)

type ConfigType struct {
	Discord     DiscordConfig
	Environment string
	Mongo       MongoConfig
}

type DiscordConfig struct {
	Activity DiscordActivityConfig
	BotID    string
	Token    string
}

type DiscordActivityConfig struct {
	Name   string
	State  string
	Type   int
	Status string
}

type MongoConfig struct {
	URL string
}
