package main

import (
	"testing"
	"unicode/utf8"
)

type TestCase struct {
	in   string
	want string
}

/* func TestReverse(t *testing.T) {
	testcases := []TestCase{
		{"The quick brown fox jumped over the lazy dog", "god yzal eht revo depmuj xof nworb kciuq ehT"},
		{"", ""},
		{"a", "a"},
		{"!12345", "54321!"},
	}
	for _, c := range testcases {
		got := ReverseString(c.in)
		if got != c.want {
			t.Errorf("ReverseString(%q) == %q, want %q", c.in, got, c.want)
		}
	}
} */

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", "", "a", "!12345"}
	//Seed
	for _, c := range testcases {
		f.Add(c)
	}

	//Fuzz
	f.Fuzz(func(t *testing.T, orig string) {
		rev, revErr := ReverseString(orig)
		doubleRev, doubleRevErr := ReverseString(rev)

		if revErr != nil || doubleRevErr != nil {
			t.Skip()
		}

		if orig != doubleRev {
			t.Errorf("Before : %q, after : %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid utf8 string: %q", rev)
		}
	})

}
