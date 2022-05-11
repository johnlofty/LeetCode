package main

import (
	"testing"
)

var tests = []struct {
	num  int
	want bool
}{
	{121, true},
	{-121, false},
	{10, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range tests {
		ans := isPalindrome(tt.num)
		if ans != tt.want {
			t.Errorf("got %t want %t", ans, tt.want)
		}
	}
}

func TestIsPalindromeStr(t *testing.T) {
	for _, tt := range tests {
		ans := isPalindromeStr(tt.num)
		if ans != tt.want {
			t.Errorf("got %t want %t", ans, tt.want)
		}
	}
}

func TestIsPalindromeBetter(t *testing.T) {
	for _, tt := range tests {
		ans := isPalindromeBetter(tt.num)
		if ans == tt.want {
			t.Errorf("got %t want %t", ans, tt.want)
		}
    }
}
