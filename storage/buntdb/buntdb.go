package buntdb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/crywolf/toyrobot/storage"
	"github.com/tidwall/buntdb"
)

type dbStorage struct {
	db *buntdb.DB
}

func NewStorage() *dbStorage {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		log.Fatal(err)
	}

	d := &dbStorage{
		db: db,
	}
	d.Reset(err, d)
	return d
}

func (s *dbStorage) Close() {
	s.db.Close()
}

func (s *dbStorage) Position() (storage.Point, error) {
	valX, err := s.get("X")
	if err != nil {
		return storage.Point{}, err
	}

	valY, err := s.get("Y")
	if err != nil {
		return storage.Point{}, err
	}

	x, err := strconv.Atoi(valX)
	if err != nil {
		return storage.Point{}, fmt.Errorf("buntdb cannot convert DB string %s to int: %v", valX, err)
	}

	y, err := strconv.Atoi(valY)
	if err != nil {
		return storage.Point{}, fmt.Errorf("buntdb cannot convert DB string %s to int: %v", valY, err)
	}

	p := storage.Point{X: x, Y: y}
	return p, nil
}

func (s *dbStorage) SetPosition(p storage.Point) error {
	err := s.set("X", strconv.Itoa(p.X))
	if err != nil {
		return err
	}

	err = s.set("Y", strconv.Itoa(p.Y))
	if err != nil {
		return err
	}

	return nil
}

func (s *dbStorage) Direction() (storage.Direction, error) {
	val, err := s.get("Direction")
	if err != nil {
		return storage.NORTH, err
	}

	var direc storage.Direction
	err = direc.FromString(val)
	if err != nil {
		return storage.NORTH, err
	}

	return direc, nil
}

func (s *dbStorage) SetDirection(d storage.Direction) error {
	err := s.set("Direction", d.String())
	if err != nil {
		return err
	}
	return nil
}

func (s *dbStorage) String() string {
	position, err := s.Position()
	if err != nil {
		return err.Error()
	}

	direc, err := s.Direction()
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%d,%d,%s", position.X, position.Y, direc)
}

func (s *dbStorage) Reset(err error, d *dbStorage) {
	err = d.set("X", "0")
	if err != nil {
		log.Fatal(err)
	}
	err = d.set("Y", "0")
	if err != nil {
		log.Fatal(err)
	}
	err = d.set("Direction", "NORTH")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *dbStorage) get(key string) (string, error) {
	var err error
	var val string
	err = s.db.View(func(tx *buntdb.Tx) error {
		val, err = tx.Get(key)
		return err
	})
	if err != nil {
		err = fmt.Errorf("buntdb: %s %v\n", key, err)
		return "", err
	}
	return val, nil
}

func (s *dbStorage) set(key, val string) error {
	err := s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, val, nil)
		return err
	})
	return err
}
