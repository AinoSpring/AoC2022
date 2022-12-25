package main

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	type args struct {
		opponent Symbol
		player   string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "WIN",
			args: args{
				opponent: ROCK,
				player:   "Z",
			},
			want: Game{opponent: ROCK, player: PAPER},
		},
		{
			name: "LOSE",
			args: args{
				opponent: PAPER,
				player:   "X",
			},
			want: Game{opponent: PAPER, player: ROCK},
		},
		{
			name: "DRAW",
			args: args{
				opponent: SCISSORS,
				player:   "Y",
			},
			want: Game{opponent: SCISSORS, player: SCISSORS},
		},
		{
			name: "WIN LONGER THAN ARRAY",
			args: args{
				opponent: SCISSORS,
				player:   "Z",
			},
			want: Game{opponent: SCISSORS, player: ROCK},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGame(tt.args.opponent, tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
