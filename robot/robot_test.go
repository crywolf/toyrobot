package robot

import (
	"bytes"
	"strings"
	"testing"

	"github.com/crywolf/toyrobot/storage/memory"
)

func TestProcessCommands(t *testing.T) {
	tests := []struct {
		name     string
		commands []Command
		want     string
		wantErr  bool
	}{
		{
			"Skips commands until PLACE command",
			[]Command{
				{"move", []string{}},
			},
			"",
			false,
		},
		{
			"Returns error if command name is illegal",
			[]Command{
				{"place", []string{"1", "1", "east"}},
				{"errcmd", []string{}},
			},
			"",
			true,
		},
		{
			"Returns error if Direction is illegal",
			[]Command{
				{"place", []string{"2", "3", "nonsense"}},
			},
			"",
			true,
		},
		{
			"Returns error if Position is illegal",
			[]Command{
				{"place", []string{"N", "3", "south"}},
			},
			"",
			true,
		},
		{
			"PLACE command",
			[]Command{
				{"place", []string{"1", "3", "east"}},
			},
			"1,3,EAST",
			false,
		},
		{
			"PLACE command returns error if Position is outside the table",
			[]Command{
				{"place", []string{"12", "3", "west"}},
			},
			"",
			true,
		},
		{
			"MOVE command",
			[]Command{
				{"place", []string{"1", "1", "east"}},
				{"move", []string{}},
			},
			"2,1,EAST",
			false,
		},
		{
			"LEFT command",
			[]Command{
				{"place", []string{"1", "1", "east"}},
				{"left", []string{}},
			},
			"1,1,NORTH",
			false,
		},
		{
			"RIGHT command",
			[]Command{
				{"place", []string{"1", "1", "south"}},
				{"right", []string{}},
			},
			"1,1,WEST",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := memory.NewStorage()
			r := NewRobot(storage, &bytes.Buffer{})

			if err := r.ProcessCommands(tt.commands); (err != nil) != tt.wantErr {
				t.Errorf("ProcessCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if position := storage.String(); tt.want != "" && position != tt.want {
				t.Errorf("ProcessCommands() = %v, want %v", position, tt.want)
			}
		})
	}

	t.Run("REPORT command writes position to output writer", func(t *testing.T) {
		storage := memory.NewStorage()
		output := &bytes.Buffer{}
		r := NewRobot(storage, output)

		commands := []Command{
			{"place", []string{"1", "2", "west"}},
			{"report", []string{}},
		}
		want := "-> position: 1,2,WEST"

		err := r.ProcessCommands(commands)
		if err != nil {
			t.Fatal(err)
		}

		if got := strings.TrimSpace(output.String()); got != "" && got != want {
			t.Errorf("ProcessCommands() writes %v to output, want %v", got, want)
		}
	})
}
