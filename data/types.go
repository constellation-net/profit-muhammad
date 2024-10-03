package data

type Scoreboard struct {
	ID     string         `bson:"_id"`
	Scores map[string]int `bson:"scores"`
}
