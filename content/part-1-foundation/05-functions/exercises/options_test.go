package exercises

import "testing"

func TestNewLogger(t *testing.T) {
	t.Skip("Chapter 05 exercise: implement the Logger options, then delete this Skip")

	cases := []struct {
		name string
		opts []Option
		want Logger
	}{
		{
			"defaults",
			nil,
			Logger{Prefix: "[LOG]", Level: "info", UseColor: false},
		},
		{
			"prefix and level",
			[]Option{WithPrefix("[APP]"), WithLevel("debug")},
			Logger{Prefix: "[APP]", Level: "debug", UseColor: false},
		},
		{
			"all options",
			[]Option{WithPrefix("[API]"), WithLevel("warn"), WithColor(true)},
			Logger{Prefix: "[API]", Level: "warn", UseColor: true},
		},
		{
			"last option wins",
			[]Option{WithLevel("debug"), WithLevel("error")},
			Logger{Prefix: "[LOG]", Level: "error", UseColor: false},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewLogger(tc.opts...)
			if got == nil {
				t.Fatal("NewLogger returned nil")
			}
			if *got != tc.want {
				t.Errorf("NewLogger(...) = %+v, want %+v", *got, tc.want)
			}
		})
	}
}
