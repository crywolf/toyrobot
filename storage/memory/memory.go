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

func (s *mstorage) Position() (storage.Point, error) {
	return s.position, nil
}

func (s *mstorage) SetPosition(p storage.Point) error {
	s.position = p
	return nil
}

func (s *mstorage) Direction() (storage.Direction, error) {
	return s.direction, nil
}

func (s *mstorage) SetDirection(d storage.Direction) error {
	s.direction = d
	return nil
}

func (s *mstorage) String() string {
	return fmt.Sprintf("%d,%d,%s", s.position.X, s.position.Y, s.direction)
}
