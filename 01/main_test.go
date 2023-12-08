package main

import (
	"testing"
)

func Test_returnNumbers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ThreeNumbers",
			args: args{
				input: "123",
			},
			want: "123",
		},
		{
			name: "TwoNumbersAndAChar",
			args: args{
				input: "1a23",
			},
			want: "123",
		},
		{
			name: "OnlyOneChar",
			args: args{
				input: "a2b",
			},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := returnNumbers(tt.args.input); got != tt.want {
				t.Errorf("returnNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTwoNumbers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lotsOfDifferentNumbers",
			args: args{
				input: "12412",
			},
			want: 12,
		},
		{
			name: "dupplicateOne",
			args: args{
				input: "1",
			},
			want: 11,
		},
		{
			name: "noNumber",
			args: args{
				input: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTwoNumbers(tt.args.input); got != tt.want {
				t.Errorf("getTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapWrittenToNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "oneTo1",
			args: args{
				input: "oneatwo",
			},
			want: "one1oneatwo2two",
		}, {
			name: "eightwo",
			args: args{
				input: "eightwo",
			},
			want: "eight8eightwo2two", // I hate it
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapWrittenToNumber(tt.args.input); got != tt.want {
				t.Errorf("mapWrittenToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
