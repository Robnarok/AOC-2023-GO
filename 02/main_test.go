package main

import (
	"reflect"
	"testing"
)

func Test_parseGames(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "firstLine",
			args: args{
				input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: []string{" 3 blue, 4 red", " 1 red, 2 green, 6 blue", " 2 green"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseGames(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSet(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Cubes
	}{
		{
			name: "3blue4red",
			args: args{
				input: " 3 blue, 4 red",
			},
			want: Cubes{
				red:   4,
				blue:  3,
				green: 0,
			},
		},
		{
			name: "1red2green6blue",
			args: args{
				input: " 1 red, 2 green, 6 blue",
			},
			want: Cubes{
				red:   1,
				blue:  6,
				green: 2,
			},
		}, {
			name: "2green",
			args: args{
				input: " 2 green",
			},
			want: Cubes{
				red:   0,
				blue:  0,
				green: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSet(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
