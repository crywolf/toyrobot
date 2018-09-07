package robot

import "io"

type Storage interface {
	Place(pos []string) error
	Step() error
	Rotate(direction string) error
	Report(w io.Writer) error
}
