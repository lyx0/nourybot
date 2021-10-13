package handlers

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/commands"
	config "github.com/lyx0/nourybot/pkg/config"
	log "github.com/sirupsen/logrus"
)

// HandleTwitchMessage takes in a twitch.Privatemessage and
// *twitch.Client and has the logic to decide if the provided
// PrivateMessage is a command or not and passes it on accordingly.
// Typical twitch message tags https://paste.ivr.fi/nopiradodo.lua
func HandleTwitchMessage(message twitch.PrivateMessage, client *twitch.Client) {
	log.Info("fn HandleTwitchMessage")
	log.Info(message)

	// roomId is the Twitch UserID of the channel the message
	// was sent in.
	roomId := message.Tags["room-id"]

	// The message has no room-id so something went wrong.
	if roomId == "" {
		log.Errorf("Missing room-id in message tag", roomId)
		return
	}

	// Message was sent from the Bot. Don't act on it
	// so that we don't repeat ourself.
	if message.Tags["user-id"] == config.LoadConfig().BotUserId {
		return
	}

	if len(message.Message) >= 2 {
		if message.Message[:2] == "()" {
			commands.HandleCommand(message, client)
		}
	}

}
