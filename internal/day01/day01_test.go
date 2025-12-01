package day01

import "testing"

func TestPart1_Example(t *testing.T) {
	input := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	got, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}

	const want = 3
	if got != want {
		t.Fatalf("Part1() = %d, want %d", got, want)
	}
}

func TestPart1_BasicWrapping(t *testing.T) {
	input := []string{
		"R50", // 50 -> 0
		"L1",  // 0 -> 99
		"R1",  // 99 -> 0
	}

	got, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}
	if got != 2 {
		t.Fatalf("Part1() = %d, want %d", got, 2)
	}
}
