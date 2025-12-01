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

func TestPart2_Example(t *testing.T) {
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

	got, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}

	const want = 6
	if got != want {
		t.Fatalf("Part2() = %d, want %d", got, want)
	}
}

func TestPart2_LargeRotationR1000(t *testing.T) {
	// Starting from 50, an R1000 rotation should land on 0 ten times
	// before returning to 50.
	input := []string{"R1000"}

	got, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}

	const want = 10
	if got != want {
		t.Fatalf("Part2() = %d, want %d", got, want)
	}
}
