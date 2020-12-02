package main

import "testing"

func Test_programPartOne(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Should be 2",  args{[]string{"1-3-a-abcde", "1-3-b-cdefg", "2-9-c-ccccccccc"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := programPartOne(tt.args.input); got != tt.want {
				t.Errorf("programPartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_programPartTwo(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Should be 1", args{[]string{"1-3-a-abcde", "1-3-b-cdefg", "2-9-c-ccccccccc"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := programPartTwo(tt.args.input); got != tt.want {
				t.Errorf("programPartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
