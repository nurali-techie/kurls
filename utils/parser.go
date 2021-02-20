package utils

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/nurali-techie/kurls/commands"
)

func GetCommand(args []string) commands.Command {
	if len(args) == 0 {
		return &commands.Help{}
	}

	var cmdName string
	if len(args) > 0 {
		cmdName = args[0]
	}

	switch cmdName {
	case "ls":
		return parseList(args[1:])
	case "add":
		return parseAdd(args[1:])
	case "get":
		return parseGet(args[1:])
	default:
		return &commands.Help{}
	}
}

func parseList(args []string) commands.Command {
	var filter string
	if len(args) > 0 {
		filter = args[0]
	}
	return &commands.List{
		Filter: filter,
	}
}

func parseAdd(args []string) commands.Command {
	var key string
	if len(args) < 1 {
		fmt.Println("error: key missing for 'add' command")
		os.Exit(1)
	}
	key = args[0]
	curl, err := clipboard.ReadAll()
	if err != nil {
		fmt.Printf("error: reading curl from clipboard, %v\n", err)
		os.Exit(1)
	}
	return commands.NewAdd(key, curl)
}

func parseGet(args []string) commands.Command {
	if len(args) < 1 {
		fmt.Println("error: missing key name")
		os.Exit(1)
	}
	return &commands.Get{
		Key: args[0],
	}
}
