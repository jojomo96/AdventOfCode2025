package day02

import (
	"fmt"
	"strconv"
	"strings"
)

// Part1 computes the sum of all invalid IDs in the given ranges using the Part 1 definition:
// a number whose decimal representation is exactly some sequence of digits repeated twice.
func Part1(lines []string) (int64, error) {
	var sum int64
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ranges, err := parseRanges(line)
		if err != nil {
			return 0, fmt.Errorf("parse ranges: %w", err)
		}
		for _, r := range ranges {
			sum += sumInvalidInRange(r.start, r.end)
		}
	}
	return sum, nil
}

// Part2 computes the sum of all invalid IDs using the Part 2 definition:
// a number made only of some sequence of digits repeated at least twice.
func Part2(lines []string) (int64, error) {
	var sum int64
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ranges, err := parseRanges(line)
		if err != nil {
			return 0, fmt.Errorf("parse ranges: %w", err)
		}
		for _, r := range ranges {
			sum += sumInvalidInRangePart2(r.start, r.end)
		}
	}
	return sum, nil
}

// rangeInclusive represents an inclusive [start, end] range.
type rangeInclusive struct {
	start int64
	end   int64
}

// parseRanges parses a single line like "11-22,95-115,..." into ranges.
func parseRanges(s string) ([]rangeInclusive, error) {
	parts := strings.Split(s, ",")
	res := make([]rangeInclusive, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		dash := strings.IndexByte(p, '-')
		if dash <= 0 || dash == len(p)-1 {
			return nil, fmt.Errorf("invalid range %q", p)
		}
		loStr := p[:dash]
		hiStr := p[dash+1:]
		lo, err := strconv.ParseInt(loStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse low %q: %w", loStr, err)
		}
		hi, err := strconv.ParseInt(hiStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse high %q: %w", hiStr, err)
		}
		if hi < lo {
			return nil, fmt.Errorf("range %q has hi < lo", p)
		}
		res = append(res, rangeInclusive{start: lo, end: hi})
	}
	return res, nil
}

// sumInvalidInRange returns the sum of all invalid IDs in [lo, hi] for Part 1.
// An invalid ID is a number consisting of some sequence of digits repeated exactly twice.
func sumInvalidInRange(lo, hi int64) int64 {
	var sum int64
	for n := lo; n <= hi; n++ {
		if isRepeatedTwice(n) {
			sum += n
		}
	}
	return sum
}

// sumInvalidInRangePart2 returns the sum of all invalid IDs in [lo, hi] for Part 2.
// An invalid ID is a number consisting only of some sequence of digits repeated at least twice.
func sumInvalidInRangePart2(lo, hi int64) int64 {
	var sum int64
	for n := lo; n <= hi; n++ {
		if isRepeatedPattern(n) {
			sum += n
		}
	}
	return sum
}

// isRepeatedTwice reports whether n's decimal representation is XY where X==Y and no leading zeros.
// This is the Part 1 predicate: exactly two repeats.
func isRepeatedTwice(n int64) bool {
	if n < 10 {
		return false
	}

	s := strconv.FormatInt(n, 10)
	if len(s)%2 != 0 {
		return false
	}

	mid := len(s) / 2
	first := s[:mid]
	second := s[mid:]

	return first == second
}

// isRepeatedPattern reports whether n's decimal representation is some block repeated
// at least twice (Part 2). No leading zeros are possible because n is a normal integer.
func isRepeatedPattern(n int64) bool {
	if n < 10 {
		return false
	}

	s := strconv.FormatInt(n, 10)
	L := len(s)

	// Try all possible block lengths from 1 up to L/2.
	for blockLen := 1; blockLen*2 <= L; blockLen++ {
		if L%blockLen != 0 {
			continue
		}
		// There will be k = L / blockLen repetitions; since blockLen*2 <= L,
		// k is at least 2 here.
		block := s[:blockLen]
		ok := true
		for i := blockLen; i < L; i += blockLen {
			if s[i:i+blockLen] != block {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	return false
}
