package storage

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type direction int

const (
	NORTH direction = iota
	EAST
	SOUTH
	WEST
)

func (d direction) String() string {
	names := [...]string{
		"NORTH",
		"EAST",
		"SOUTH",
		"WEST"}

	if d < NORTH || d > WEST {
		return "Unknown"
	}

	return names[d]
}

func (d *direction) FromString(s string) {
	switch strings.ToLower(s) {
	case "north":
		*d = NORTH
	case "east":
		*d = EAST
	case "south":
		*d = SOUTH
	case "west":
		*d = WEST
	}
}

type point struct {
	X, Y int
}

type storage struct {
	Position  point
	Direction direction
}

func NewStorage() *storage {
	return &storage{
		Position:  point{},
		Direction: NORTH,
	}
}

func (s *storage) Place(pos []string) error {
	x, err := strconv.Atoi(pos[0])
	if err != nil {
		return fmt.Errorf("cannot convert string %s to int: %v", pos[0], err)
	}

	y, err := strconv.Atoi(pos[1])
	if err != nil {
		return fmt.Errorf("cannot convert string %s to int: %v", pos[1], err)
	}

	if x < 0 || x > 4 || y < 0 || y > 4 {
		return fmt.Errorf("illegal placement")
	}
	s.Position.X = x
	s.Position.Y = y
	s.Direction.FromString(pos[2])
	return nil
}

func (s *storage) Step() error {
	switch s.Direction {
	case NORTH:
		if s.Position.Y < 4 {
			s.Position.Y++
		}
	case SOUTH:
		if s.Position.Y > 0 {
			s.Position.Y--
		}
	case EAST:
		if s.Position.X < 4 {
			s.Position.X++
		}
	case WEST:
		if s.Position.X > 0 {
			s.Position.X--
		}
	}
	return nil
}

func (s *storage) Rotate(direction string) error {
	if direction == "left" {
		if s.Direction == 0 {
			s.Direction += 3
		} else {
			s.Direction -= 1
		}
	} else if direction == "right" {
		if s.Direction == 3 {
			s.Direction -= 3
		} else {
			s.Direction += 1
		}
	}
	return nil
}

func (s *storage) Report(w io.Writer) error {
	fmt.Fprintln(w, "-> position:", s)
	return nil
}

func (s *storage) String() string {
	return fmt.Sprintf("%d,%d,%s", s.Position.X, s.Position.Y, s.Direction)
}
