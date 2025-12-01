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
