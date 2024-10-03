package data

import (
	"context"

	"github.com/constellation-net/profit-muhammad/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ScoreboardCollection = Client.Database("profit-muhammad").Collection("scoreboards")
)

func NewScoreboard(id string) Scoreboard {
	scoreboard := Scoreboard{
		ID:     id,
		Scores: map[string]int{},
	}
	_, err := ScoreboardCollection.InsertOne(context.TODO(), scoreboard)
	if err != nil {
		log.Error(err, "SCOREBOARD_NEW", true)
	}

	return scoreboard
}

func GetScoreboard(id string) Scoreboard {
	var result Scoreboard
	filter := bson.D{{Key: "_id", Value: id}}

	err := ScoreboardCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return NewScoreboard(id)
	} else if err != nil {
		log.Error(err, "SCOREBOARD_FIND", true)
	}

	return result
}

func GetUserScore(id string, userID string) int {
	scoreboard := GetScoreboard(id)
	return scoreboard.Scores[userID]
}

func SetUserScore(id string, userID string, score int) {
	scoreboard := GetScoreboard(id)
	scoreboard.Scores[userID] = score
	_, err := ScoreboardCollection.ReplaceOne(context.TODO(), bson.D{{Key: "_id", Value: id}}, scoreboard)
	if err != nil {
		log.Error(err, "SCOREBOARD_SET", true)
	}
}

func IncrementUserScore(id string, userID string) int {
	score := GetUserScore(id, userID) + 1
	SetUserScore(id, userID, score)
	return score
}
