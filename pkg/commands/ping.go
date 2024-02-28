package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/pkg/common"
	"github.com/lyx0/nourybot/pkg/humanize"
)

// Ping returns information about the bot.
func Ping(env string) string {
	botUptime := humanize.Time(common.GetUptime())
	commandsUsed := common.GetCommandsUsed()
	commit := common.GetVersion()

	return fmt.Sprintf("Pong! :) Commands used: %v, Last restart: %v, Running on commit: %v, Env: %v", commandsUsed, botUptime, commit, env)
}
