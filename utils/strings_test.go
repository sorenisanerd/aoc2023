package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestExtractInts(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"", []int{}},
		{"12 34 45", []int{12, 34, 45}},
		{"12x34y45", []int{12, 34, 45}},
		{"asdfasdf12x34y45", []int{12, 34, 45}},
		{"12      \t34\n45x", []int{12, 34, 45}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.DeepEqual(t, tt.want, ExtractInts(tt.name))
		})
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 string
	}{
		{"1234blah", 1234, "blah"},
		{"1234 blah", 1234, " blah"},
		{"x1234 blah", 0, "x1234 blah"},
		{"1234", 1234, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Atoi(tt.name)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestMustAtoi(t *testing.T) {
	tests := []struct {
		name        string
		want        int
		shouldPanic bool
	}{
		{"1234", 1234, false},
		{"-1234", -1234, false},
		{"1234dsaffafs", 1234, true},
		{"foo1234", 0, true},
		{"", 0, true},
		{"-1234dsaffafs", 1234, true},
		{"-foo1234", 0, true},
		{"-", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				didPanic := recover() != nil
				if didPanic != tt.shouldPanic {
					t.Errorf("didPanic = %v, wantPanic = %v", didPanic, tt.shouldPanic)
				}
			}()
			assert.Equal(t, tt.want, MustAtoi(tt.name))
		})
	}
}

func TestReplaceAtIndex(t *testing.T) {
	type args struct {
		in string
		i  int
		r  rune
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantPanic bool
	}{
		{`("", 0, 'a')`, args{"", 0, 'a'}, "who cares? It's going to panic.", true},
		{"fao->foo", args{"fao", 1, 'o'}, "foo", false},
		{"foa->foo", args{"foa", 2, 'o'}, "foo", false},
		{"foa->foo", args{"foa", -1, 'o'}, "foo", false},
		{"foa->foo", args{"foa", 3, 'o'}, "foo", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				didPanic := recover() != nil
				if didPanic != tt.wantPanic {
					t.Errorf("didPanic = %v, wantPanic = %v", didPanic, tt.wantPanic)
				}
			}()
			assert.Equal(t, tt.want, ReplaceAtIndex(tt.args.in, tt.args.i, tt.args.r))
		})
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"", []string{}},
		{"foo bar", []string{"foo", "bar"}},
		{"foo      bar", []string{"foo", "bar"}},
		{"     foo      bar", []string{"foo", "bar"}},
		{"     foo      bar", []string{"foo", "bar"}},
		{"  asdf\nqwer\tzxcv   foo      bar", []string{"asdf", "qwer", "zxcv", "foo", "bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.DeepEqual(t, tt.want, Words(tt.name))
		})
	}
}
