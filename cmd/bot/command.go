package main

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/commands"
	"github.com/lyx0/nourybot/pkg/common"
	"github.com/sirupsen/logrus"
)

func handleCommand(message twitch.PrivateMessage, tc *twitch.Client) {
	logrus.Info("[COMMAND HANDLER]", message)
	logrus.Info(message)

	// commandName is the actual name of the command without the prefix.
	// e.g. `()ping` would be `ping`.
	commandName := strings.ToLower(strings.SplitN(message.Message, " ", 3)[0][2:])
	logrus.Info(commandName)

	// cmdParams are additional command parameters.
	// e.g. `()weather san antonio`
	// cmdParam[0] is `san` and cmdParam[1] = `antonio`.
	//
	// Since Twitch messages are at most 500 characters I use a
	// maximum count of 500+10 just to be safe.
	// https://discuss.dev.twitch.tv/t/missing-client-side-message-length-check/21316
	cmdParams := strings.SplitN(message.Message, " ", 500)
	_ = cmdParams

	// msgLen is the amount of words in a message without the prefix.
	// Useful to check if enough cmdParams are provided.
	msgLen := len(strings.SplitN(message.Message, " ", -2))
	logrus.Info(msgLen)

	// target is the channelname the message originated from and
	// where we are responding.
	target := message.Channel

	switch commandName {
	case "":
		if msgLen == 1 {
			common.Send(target, "xd", tc)
			return
		}
	case "echo":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", tc)
			return
		} else {
			commands.Echo(target, message.Message[7:len(message.Message)], tc)
			// bot.Send(target, message.Message[7:len(message.Message)])
			return
		}
	}
}
