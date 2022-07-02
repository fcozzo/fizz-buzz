package main

import "testing"

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Fizz",
			args: args{number: 6},
			want: "Fizz",
		},
		{
			name: "Test Buzz",
			args: args{number: 25},
			want: "Buzz",
		},
		{
			name: "Test Fizz Buzz",
			args: args{number: 75},
			want: "Fizz Buzz",
		},
		{
			name: "Test Number",
			args: args{number: 19},
			want: "19",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.number); got != tt.want {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
