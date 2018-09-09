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

func ProcessCommands(commands []Command, storage storage.Storage, output io.Writer) error {
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
			err = place(storage, command.args)
		case "move":
			err = step(storage)
		case "left", "right":
			err = rotate(storage, command.cmd)
		case "report":
			err = report(storage, output)
		default:
			err = fmt.Errorf("unknown command '%s'", command.cmd)
		}
		if err != nil {
			break
		}
	}
	return err
}

func place(s storage.Storage, pos []string) error {
	position := s.Position()
	direction := s.Direction()

	x, err := strconv.Atoi(pos[0])
	if err != nil {
		return fmt.Errorf("cannot convert string %s to int: %v", pos[0], err)
	}

	y, err := strconv.Atoi(pos[1])
	if err != nil {
		return fmt.Errorf("cannot convert string %s to int: %v", pos[1], err)
	}

	if x < 0 || x > maxX || y < 0 || y > maxY {
		return fmt.Errorf("illegal placement (illegal position[%d,%d])", x, y)
	}

	position.X = x
	position.Y = y

	err = direction.FromString(pos[2])
	if err != nil {
		return fmt.Errorf("illegal placement (%v)", err)
	}

	s.SetPosition(position)
	s.SetDirection(direction)
	return nil
}

func step(s storage.Storage) error {
	pos := s.Position()
	direc := s.Direction()

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

	s.SetPosition(pos)
	s.SetDirection(direc)
	return nil
}

func rotate(s storage.Storage, direction string) error {
	direc := s.Direction()

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

	s.SetDirection(direc)
	return nil
}

func report(s storage.Storage, w io.Writer) error {
	fmt.Fprintln(w, "-> position:", s)
	return nil
}
