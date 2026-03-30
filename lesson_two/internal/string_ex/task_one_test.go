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
		{"case 3", "45", ""},
		{"case 3", "aaa10b", ""},
		{"case 3", "aaa0b", "aab"},
		{"case 4", "", ""},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PackString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
