package storage

import (
	"fmt"
	"strings"
)

type Storage interface {
	Position() point
	SetPosition(point)
	Direction() direction
	SetDirection(direction)
	fmt.Stringer
}

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

func (d *direction) FromString(s string) error {
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

type point struct {
	X, Y int
}

type storage struct {
	position  point
	direction direction
}

func NewStorage() *storage {
	return &storage{
		position:  point{},
		direction: NORTH,
	}
}

func (s *storage) Position() point {
	return s.position
}

func (s *storage) SetPosition(p point) {
	s.position = p
}

func (s *storage) Direction() direction {
	return s.direction
}

func (s *storage) SetDirection(d direction) {
	s.direction = d
}

func (s *storage) String() string {
	return fmt.Sprintf("%d,%d,%s", s.position.X, s.position.Y, s.direction)
}
