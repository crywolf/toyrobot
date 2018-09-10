package robot

import (
	"fmt"
	"io"
	"strconv"

	"github.com/crywolf/toyrobot/storage"
)

// playground size (5x5 units, indices start with 0)
const (
	maxX = 4
	maxY = 4
)

type robot struct {
	storage storage.Storage
	output  io.Writer
}

func NewRobot(storage storage.Storage, output io.Writer) *robot {
	return &robot{
		storage: storage,
		output:  output,
	}
}

func (r *robot) ProcessCommands(commands []Command) error {
	var err error
	skip := true
	for _, command := range commands {
		// skip commands before first PLACE command
		if skip && command.cmd != "place" {
			continue
		} else {
			skip = false
		}

		switch command.cmd {
		case "place":
			err = r.place(command.args)
		case "move":
			err = r.step()
		case "left", "right":
			err = r.rotate(command.cmd)
		case "report":
			err = r.report()
		default:
			err = fmt.Errorf("robot: unknown command '%s'", command.cmd)
		}
		if err != nil {
			break
		}
	}
	return err
}

func (r *robot) place(pos []string) error {
	position, err := r.storage.Position()
	if err != nil {
		return err
	}

	direction, err := r.storage.Direction()
	if err != nil {
		return err
	}

	x, err := strconv.Atoi(pos[0])
	if err != nil {
		return fmt.Errorf("robot: cannot convert string %s to int: %v", pos[0], err)
	}

	y, err := strconv.Atoi(pos[1])
	if err != nil {
		return fmt.Errorf("robot: cannot convert string %s to int: %v", pos[1], err)
	}

	if x < 0 || x > maxX || y < 0 || y > maxY {
		return fmt.Errorf("robot: illegal placement (illegal position[%d,%d])", x, y)
	}

	position.X = x
	position.Y = y

	err = direction.FromString(pos[2])
	if err != nil {
		return fmt.Errorf("robot: illegal placement (%v)", err)
	}

	err = r.storage.SetPosition(position)
	if err != nil {
		return err
	}

	err = r.storage.SetDirection(direction)
	if err != nil {
		return err
	}

	return nil
}

func (r *robot) step() error {
	pos, err := r.storage.Position()
	if err != nil {
		return err
	}

	direc, err := r.storage.Direction()
	if err != nil {
		return err
	}

	switch direc {
	case storage.NORTH:
		if pos.Y < maxY {
			pos.Y++
		}
	case storage.SOUTH:
		if pos.Y > 0 {
			pos.Y--
		}
	case storage.EAST:
		if pos.X < maxX {
			pos.X++
		}
	case storage.WEST:
		if pos.X > 0 {
			pos.X--
		}
	}

	err = r.storage.SetPosition(pos)
	if err != nil {
		return err
	}

	err = r.storage.SetDirection(direc)
	if err != nil {
		return err
	}

	return nil
}

func (r *robot) rotate(direction string) error {
	direc, err := r.storage.Direction()
	if err != nil {
		return err
	}

	if direction == "left" {
		if direc == 0 {
			direc += 3
		} else {
			direc -= 1
		}
	} else if direction == "right" {
		if direc == 3 {
			direc -= 3
		} else {
			direc += 1
		}
	}

	err = r.storage.SetDirection(direc)
	if err != nil {
		return err
	}

	return nil
}

func (r *robot) report() error {
	_, err := fmt.Fprintln(r.output, "-> position:", r.storage)
	if err != nil {
		return fmt.Errorf("robot could not write to output stream: %v", err)
	}
	return nil
}
