package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	URI       string `json:"uri"`
	Timeout   string `json:"timeout"`
	Buffer    string `json:"buffer"`
	Throttle  string `json:"throttle"`
	Heartbeat string `json:"heartbeat"`
	Auth      Auth   `json:"auth"`
	Output    Output `json:"output"`
	Data      Data   `json:"data"`
}

type Auth struct {
	Token string `json:"token"`
}

type Output struct {
	PrecisionTime   string `json:"precision_time"`
	PrecisionPos    string `json:"precision_position"`
	PrecisionVector string `json:"precision_vector"`
}

type Data struct {
	Provider             string `json:"provider"`
	Map                  string `json:"map"`
	MapRoundWins         string `json:"map_round_wins"`
	Round                string `json:"round"`
	PlayerID             string `json:"player_id"`
	PlayerState          string `json:"player_state"`
	PlayerWeapons        string `json:"player_weapons"`
	PlayerMatchStats     string `json:"player_match_stats"`
	AllGrenades          string `json:"allgrenades"`
	AllPlayersID         string `json:"allplayers_id"`
	AllPlayersMatchStats string `json:"allplayers_match_stats"`
	AllPlayersPosition   string `json:"allplayers_position"`
	AllPlayersState      string `json:"allplayers_state"`
	AllPlayersWeapons    string `json:"allplayers_weapons"`
	Bomb                 string `json:"bomb"`
	PhaseCountdowns      string `json:"phase_countdowns"`
	PlayerPosition       string `json:"player_position"`
}

var dataJson = []byte(`[{
  "uri": "http://127.0.0.1:5000",
  "timeout": "5.0",
  "buffer": "0.1",
  "throttle": "0.1",
  "heartbeat": "30.0",
  "auth": {
    "token": "Q79v5tcxVQ8u"
  },
  "output": {
    "precision_time": "3",
    "precision_position": "1",
    "precision_vector": "3"
  },
  "data": {
    "provider": "1",
    "map": "1",
    "map_round_wins": "1",
    "round": "1",
    "player_id": "1",
    "player_state": "1",
    "player_weapons": "1",
    "player_match_stats": "1",
    "allgrenades": "1",
    "allplayers_id": "1",
    "allplayers_match_stats": "1",
    "allplayers_position": "1",
    "allplayers_state": "1",
    "allplayers_weapons": "1",
    "bomb": "1",
    "phase_countdowns": "1",
    "player_position": "1"
  }
}]`)

func main() {

	var configs []Config
	err := json.Unmarshal(dataJson, &configs)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Access the first configuration object
	config := configs[0]

	fmt.Printf("URI: %s\n", config.URI)
	fmt.Printf("Token: %s\n", config.Auth.Token)
	fmt.Printf("Precision Time: %s\n", config.Output.PrecisionTime)
	fmt.Printf("Player Position: %s\n", config.Data.PlayerPosition)
	http.HandleFunc("/post", handlePost)

	log.Println("Starting HTTP server on :5000...")
	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

//response = requests.post('http://127.0.0.1:5000',data,auth='rx5w2bXmCCWJu6')
