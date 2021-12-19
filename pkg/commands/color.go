package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lyx0/nourybot/cmd/bot"
	log "github.com/sirupsen/logrus"
)

// https://api.ivr.fi
type colorApiResponse struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	ChatColor   string `json:"chatColor"`
	Error       string `json:"error"`
}

// Userid returns the userID of a given user
func Color(username, target string, nb *bot.Bot) {
	baseUrl := "https://api.ivr.fi/twitch/resolve"

	resp, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, username))
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject colorApiResponse
	json.Unmarshal(body, &responseObject)

	// time string format 2011-05-19T00:28:28.310449Z
	// discard everything after T

	reply := fmt.Sprintf("%s color is %s",
		responseObject.DisplayName,
		responseObject.ChatColor,
	)

	// User not found
	if responseObject.Error != "" {
		nb.Send(target, "Something went wrong... FeelsBadMan")
	}
	nb.Send(target, reply)
}
