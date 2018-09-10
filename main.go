package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/crywolf/toyrobot/robot"
	"github.com/crywolf/toyrobot/storage"
	"github.com/crywolf/toyrobot/storage/memory"
)

func main() {
	db := memory.NewStorage()
	err := start(os.Args[1:], db, os.Stdout)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func start(programArgs []string, storage storage.Storage, output io.Writer) error {
	commands := prepareCommands(programArgs)
	r := robot.NewRobot(storage, output)
	return r.ProcessCommands(commands)
}

// parses commandline and creates corresponding commands
func prepareCommands(programArgs []string) []robot.Command {
	var commands []robot.Command

	argumentFound := false
	for i, arg := range programArgs {
		loweredArg := strings.ToLower(arg)
		if argumentFound {
			argumentFound = false
			continue
		}
		if loweredArg == "place" {
			placeArgs := strings.Split(programArgs[i+1], ",")
			c := robot.MakeCommand(loweredArg, placeArgs)
			commands = append(commands, c)
			argumentFound = true
		} else {
			c := robot.MakeCommand(loweredArg, nil)
			commands = append(commands, c)
		}
	}
	return commands
}
