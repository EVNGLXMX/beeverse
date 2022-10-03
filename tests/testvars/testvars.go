package testvars

import (
	"beeverse/models/mymongo"
	"time"
)

type MetaData struct {
	ServerStatus bool  `json:"server_status" bson:"server_status"`
	PlayerCount  int64 `json:"player_count" bson:"player_count"`
	Updated      bool  `json:"updated" bson:"updated"`
}

var TestGame = mymongo.Game{
	GameID:      "testgame",
	Genres:      []string{"testGenre", "testGenre2"},
	ReleaseDate: time.Date(2022, time.September, 29, 0, 0, 0, 0, time.UTC),
	CreatedOn:   time.Now(),
	UpdatedOn:   time.Now(),
}
