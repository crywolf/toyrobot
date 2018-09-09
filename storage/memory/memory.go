package memory

import (
	"fmt"

	"github.com/crywolf/toyrobot/storage"
)

type mstorage struct {
	position  storage.Point
	direction storage.Direction
}

func NewStorage() *mstorage {
	return &mstorage{
		position:  storage.Point{},
		direction: storage.NORTH,
	}
}

func (s *mstorage) Position() storage.Point {
	return s.position
}

func (s *mstorage) SetPosition(p storage.Point) {
	s.position = p
}

func (s *mstorage) Direction() storage.Direction {
	return s.direction
}

func (s *mstorage) SetDirection(d storage.Direction) {
	s.direction = d
}

func (s *mstorage) String() string {
	return fmt.Sprintf("%d,%d,%s", s.position.X, s.position.Y, s.direction)
}
