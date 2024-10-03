package data

import (
	"context"

	"github.com/constellation-net/profit-muhammad/config"
	"github.com/constellation-net/profit-muhammad/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
)

func init() {
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Config.Mongo.URL))
	if err != nil {
		log.Error(err, "MONGO_CONNECT", true)
	}

	if err := Client.Ping(context.TODO(), nil); err != nil {
		log.Error(err, "MONGO_PING", true)
	}
}

func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Error(err, "MONGO_DISCONNECT", true)
	}
}
