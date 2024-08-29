package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Player struct {
		Activity string `json:"activity"`
		State    struct {
			Health  int `json:"health"`
			Flashed int `json:"flashed"`
			Burning int `json:"burning"`
		} `json:"state"`
	} `json:"player"`
	Round struct {
		Bomb    string `json:"bomb"`
		WinTeam string `json:"win_team"`
		Phase   string `json:"phase"`
	} `json:"round"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var payload Payload
		err = json.Unmarshal(body, &payload)
		if err != nil {
			http.Error(w, "Failed to parse JSON payload", http.StatusBadRequest)
			return
		}

		// Process the payload

		fmt.Fprintf(w, "POST request processed successfully.")

	} else {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
func processPayload(p *Payload) {
	// Your logic here
	// Example: Print out the player's health

	fmt.Printf("Player Health: %d\n", p.Player.State.Health)
}
