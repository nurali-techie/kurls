package utils

import (
	"strings"

	"github.com/nurali-techie/kurls/commands"
)

func GetCommand(args []string) commands.Command {
	if len(args) == 0 {
		return nil
	}

	var cmdName string
	if len(args) > 0 {
		cmdName = args[0]
	}

	switch cmdName {
	case "list":
		var filter string
		if len(args) > 1 {
			filter = args[1]
		}
		return &commands.List{Filter: filter}
	default:
		return parseCommand(args)
	}
}

func parseCommand(args []string) commands.Command {
	if len(args) == 0 {
		return nil
	}
	if len(args) == 1 {
		return &commands.Get{
			Key: args[0],
		}
	}

	if strings.HasPrefix(args[1], "curl") {
		return commands.NewAdd(args[0], args[1])

	}

	return nil
}
