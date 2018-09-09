package robot

type Command struct {
	cmd  string
	args []string
}

func MakeCommand(cmd string, args []string) Command {
	if args == nil {
		args = make([]string, 3)
	}
	return Command{cmd, args}
}
