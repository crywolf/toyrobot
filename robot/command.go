package robot

import (
	"fmt"
	"io"
)

type Command struct {
	cmd  string
	args []string
}

func NewCommand(cmd string, args []string) Command {
	if args == nil {
		args = make([]string, 3)
	}
	return Command{cmd, args}
}

func ProcessCommands(commands []Command, storage Storage, output io.Writer) error {
	var err error
	for _, command := range commands {
		switch command.cmd {
		case "place":
			err = storage.Place(command.args)
		case "move":
			err = storage.Step()
		case "left", "right":
			err = storage.Rotate(command.cmd)
		case "report":
			err = storage.Report(output)
		default:
			err = fmt.Errorf("unknown command '%s'", command.cmd)
		}
		if err != nil {
			break
		}
	}
	return err
}
