package day01

import (
	"fmt"
	"strconv"
	"strings"
)

// Part1 computes how many times the dial points at 0 during the sequence of rotations.
//
// Contract:
//   - lines: each non-empty line is of the form `L<number>` or `R<number>`.
//   - Dial has positions 0..99 and starts at 50.
//   - Turning left decreases the position, turning right increases it, wrapping around.
//   - We count how many times after applying a rotation the dial is exactly at 0.
//   - Empty/comment lines are ignored.
func Part1(lines []string) (int, error) {
	const modulo = 100
	pos := 50
	countZero := 0

	for i, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		if len(line) < 2 {
			return 0, fmt.Errorf("line %d: too short: %q", i+1, line)
		}

		dir := line[0]
		distStr := strings.TrimSpace(line[1:])
		dist, err := strconv.Atoi(distStr)
		if err != nil {
			return 0, fmt.Errorf("line %d: invalid distance %q: %w", i+1, distStr, err)
		}

		steps := dist % modulo
		switch dir {
		case 'L', 'l':
			pos = (pos - steps) % modulo
			if pos < 0 {
				pos += modulo
			}
		case 'R', 'r':
			pos = (pos + steps) % modulo
		default:
			return 0, fmt.Errorf("line %d: invalid direction %q", i+1, string(dir))
		}

		if pos == 0 {
			countZero++
		}
	}

	return countZero, nil
}

// Part2 counts how many times the dial points at 0 on any click, including intermediate
// positions during each rotation.
//
// For each rotation of distance D:
//   - We conceptually move the dial 1 click at a time in the given direction.
//   - Each time a click lands exactly on 0, we increment the counter.
//   - Large distances are handled efficiently by decomposing steps around the ring.
func Part2(lines []string) (int, error) {
	const modulo = 100
	pos := 50
	countZero := 0

	for i, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		if len(line) < 2 {
			return 0, fmt.Errorf("line %d: too short: %q", i+1, line)
		}

		dir := line[0]
		distStr := strings.TrimSpace(line[1:])
		dist, err := strconv.Atoi(distStr)
		if err != nil {
			return 0, fmt.Errorf("line %d: invalid distance %q: %w", i+1, distStr, err)
		}

		if dist == 0 {
			continue
		}

		steps := dist % modulo
		fullLoops := dist / modulo

		// Each full loop around the dial in either direction touches 0 exactly once.
		if fullLoops > 0 {
			// But only count the 0s that are actually reached; if the current position is 0,
			// then the first click of the loop leaves 0, so we still count exactly once.
			countZero += fullLoops
		}

		// Apply the remaining steps one by one to correctly count any additional 0 hits.
		stepDelta := 0
		switch dir {
		case 'L', 'l':
			stepDelta = -1
		case 'R', 'r':
			stepDelta = 1
		default:
			return 0, fmt.Errorf("line %d: invalid direction %q", i+1, string(dir))
		}

		for s := 0; s < steps; s++ {
			pos = (pos + stepDelta) % modulo
			if pos < 0 {
				pos += modulo
			}
			if pos == 0 {
				countZero++
			}
		}
	}

	return countZero, nil
}
