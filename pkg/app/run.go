package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/liampulles/banger/commands"
)

func Run(args []string) int {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "you must provide a subCommand - valid subCommands: %s\n", subCommands())
		return 1
	}
	subCommand := commands.Registered[args[1]]
	if err := subCommand(args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		return 1
	}
	return 0
}

func subCommands() string {
	var names []string
	for name := range commands.Registered {
		names = append(names, name)
	}
	return strings.Join(names, ", ")
}
