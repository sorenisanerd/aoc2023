package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1 2", 1},
		{"12 24", 12},
		{"24 12", 12},
		{"81 450", 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ab := ExtractInts(tt.name)
			assert.Equal(t, tt.want, GCD(ab[0], ab[1]))
		})
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1 2", 2},
		{"213412 435612", 23241207036},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ints := ExtractInts(tt.name)
			assert.Equal(t, tt.want, LCM(ints[0], ints[1:]...))
		})
	}
}

func TestIsDigitString(t *testing.T) {
	tests := []struct {
		name      string
		want      bool
		wantPanic bool
	}{
		{"1", true, false},
		{"2", true, false},
		{"3", true, false},
		{"4", true, false},
		{"5", true, false},
		{"6", true, false},
		{"7", true, false},
		{"8", true, false},
		{"9", true, false},
		{"0", true, false},
		{"a", false, false},
		{"", false, true},
		{"  ", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				didPanic := recover() != nil
				if didPanic != tt.wantPanic {
					t.Errorf("didPanic = %v, wantPanic = %v", didPanic, tt.wantPanic)
				}
			}()

			assert.Equal(t, tt.want, IsDigit(tt.name))
		})
	}
}

func TestIsDigitByteAndRune(t *testing.T) {
	tests := []struct {
		arg       byte
		want      bool
		wantPanic bool
	}{
		{'1', true, false},
		{'2', true, false},
		{'3', true, false},
		{'4', true, false},
		{'5', true, false},
		{'6', true, false},
		{'7', true, false},
		{'8', true, false},
		{'9', true, false},
		{'0', true, false},
		{'a', false, false},
		{' ', false, true},
	}
	for _, tt := range tests {
		t.Run(string(tt.arg), func(t *testing.T) {
			assert.Equal(t, tt.want, IsDigit(tt.arg))
			assert.Equal(t, tt.want, IsDigit(rune(tt.arg)))
		})
	}
}
