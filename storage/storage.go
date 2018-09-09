package storage

import (
	"fmt"
	"strings"
)

type Storage interface {
	Position() (Point, error)
	SetPosition(Point) error
	Direction() (Direction, error)
	SetDirection(Direction) error
	fmt.Stringer
}

type Point struct {
	X, Y int
}

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func (d Direction) String() string {
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

func (d *Direction) FromString(s string) error {
	switch strings.ToLower(s) {
	case "north":
		*d = NORTH
	case "east":
		*d = EAST
	case "south":
		*d = SOUTH
	case "west":
		*d = WEST
	default:
		return fmt.Errorf("illegal direction string value '%s'", s)
	}
	return nil
}
