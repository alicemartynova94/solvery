package lesson_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"case 1", "a4bc2d5e", "aaaabccddddde"},
		{"case 2", "abcd", "abcd"},
		{"case 3", "3abc", ""},
		{"case 4", "45", ""},
		{"case 5", "aaa10b", ""},
		{"case 6", "aaa0b", "aab"},
		{"case 7", "", ""},
		{"case 8", "qwe\\4\\5", "qwe45"},
		{"case 9", "qwe\\45", "qwe44444"},
		{"case 10", "qwe\\\\5", "qwe\\\\\\\\\\"},
		{"case 11", "qw\\ne", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnpackString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPackString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"case 1", "aaaabccddddde", "a4bc2d5e"},
		{"case 2", "abcd", "abcd"},
		{"case 3", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PackString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
