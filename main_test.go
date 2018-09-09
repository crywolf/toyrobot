package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/crywolf/toyrobot/storage"
)

func Test_start(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		exp     string
		wantErr bool
	}{
		{
			"PLACE 0,0,NORTH MOVE REPORT",
			"PLACE 0,0,NORTH MOVE REPORT",
			"-> position: 0,1,NORTH",
			false,
		},
		{
			"PLACE 0,0,NORTH LEFT REPORT",
			"PLACE 0,0,NORTH LEFT REPORT",
			"-> position: 0,0,WEST",
			false,
		},
		{
			"PLACE 1,2,EAST MOVE MOVE LEFT MOVE REPORT",
			"PLACE 1,2,EAST MOVE MOVE LEFT MOVE REPORT",
			"-> position: 3,3,NORTH",
			false,
		},

		{
			"Try to move outside the table 1",
			"PLACE 0,0,NORTH LEFT MOVE REPORT",
			"-> position: 0,0,WEST",
			false,
		},
		{
			"Try to move outside the table 2",
			"PLACE 2,4,NORTH MOVE REPORT",
			"-> position: 2,4,NORTH",
			false,
		},
		{
			"Try to move outside the table 3",
			"PLACE 3,2,SOUTH LEFT MOVE MOVE REPORT",
			"-> position: 4,2,EAST",
			false,
		},

		{
			"Try to place outside the table 1 ends with error",
			"PLACE -1,2,SOUTH LEFT MOVE MOVE REPORT",
			"",
			true,
		},

		{
			"Skip commands before PLACE command",
			"MOVE MOVE REPORT PLACE 3,2,WEST MOVE REPORT",
			"-> position: 2,2,WEST",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := storage.NewStorage()
			var output bytes.Buffer

			args := strings.Split(tt.args, " ")
			err := start(args, storage, &output)
			if (err != nil) != tt.wantErr {
				t.Errorf("start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got := strings.TrimSpace(output.String()); got != tt.exp {
				t.Errorf("expected \"%s\", got \"%s\"", tt.exp, got)
			}
		})
	}
}
