package random

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringWithCharset(t *testing.T) {

	tests := []struct {
		length     uint
		charset    string
		wantLength uint
	}{
		{
			length:     5,
			charset:    "ABCD",
			wantLength: 5,
		},
		{
			length:     0,
			charset:    "ABCD",
			wantLength: 0,
		},
		{
			length:     5,
			charset:    "",
			wantLength: 0,
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("test with length %d and charset %v", tt.length, tt.charset)
		t.Run(name, func(t *testing.T) {
			got := StringWithCharset(tt.length, tt.charset)
			if uint(len(got)) != tt.wantLength {
				t.Fatalf("got: %v, got length: %d, want length: %d", got, len(got), tt.wantLength)
			}
			for _, c := range got {
				if !strings.Contains(tt.charset, string(c)) {
					t.Fatalf("got: %v, char: %v, charset: %v", got, c, tt.charset)
				}
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		length uint
	}{
		{
			length: 2,
		},
		{
			length: 0,
		},
		{
			length: 100,
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("test with length %d", tt.length)
		t.Run(name, func(t *testing.T) {
			got := String(tt.length)
			if uint(len(got)) != tt.length {
				t.Fatalf("got: %v length: %d, want length: %d", got, len(got), tt.length)
			}
		})
	}
}
