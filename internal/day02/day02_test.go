package day02

import "testing"

func TestIsRepeatedTwice(t *testing.T) {
	cases := []struct {
		n    int64
		want bool
	}{
		{11, true},
		{22, true},
		{55, true},
		{64, false},
		{6464, true},
		{123123, true},
		{1010, true},
		{101, false},
		{1, false},
		{1234, false},
	}
	for _, tc := range cases {
		if got := isRepeatedTwice(tc.n); got != tc.want {
			t.Fatalf("isRepeatedTwice(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}

func TestExampleSumInvalidInRange(t *testing.T) {
	// Example ranges from the problem statement; we just spot-check a couple of ranges
	if got := sumInvalidInRange(11, 22); got != 33 { // 11 + 22
		t.Fatalf("sumInvalidInRange(11,22) = %d, want 33", got)
	}

	if got := sumInvalidInRange(95, 115); got != 99 {
		t.Fatalf("sumInvalidInRange(95,115) = %d, want 99", got)
	}
}

func TestIsRepeatedPattern(t *testing.T) {
	cases := []struct {
		n    int64
		want bool
	}{
		// From problem statement examples
		{12341234, true},   // 1234 two times
		{123123123, true},  // 123 three times
		{1212121212, true}, // 12 five times
		{1111111, true},    // 1 seven times
		// From earlier invalids that should still be invalid
		{11, true}, // 1 two times
		{22, true},
		{99, true},
		{1010, true},
		{1188511885, true},
		{446446, true},
		{38593859, true},
		// Some non-patterns
		{64, false},
		{101, false},
		{1234, false},
	}

	for _, tc := range cases {
		if got := isRepeatedPattern(tc.n); got != tc.want {
			t.Fatalf("isRepeatedPattern(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}

func TestExamplePart2Totals(t *testing.T) {
	// We'll stitch together the example ranges into the format our parser expects.
	line := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224," +
		"1698522-1698528,446443-446449,38593856-38593862,565653-565659," +
		"824824821-824824827,2121212118-2121212124"

	got, err := Part2([]string{line})
	if err != nil {
		t.Fatalf("Part2(example) error = %v", err)
	}

	const want int64 = 4174379265
	if got != want {
		t.Fatalf("Part2(example) = %d, want %d", got, want)
	}
}
