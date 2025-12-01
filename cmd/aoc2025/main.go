package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"AdventOfCode2025/internal/day01"
)

func main() {
	day := flag.Int("day", 1, "day of Advent of Code 2025 to run")
	part := flag.Int("part", 1, "puzzle part to run (1 or 2)")
	inputPath := flag.String("input", "input/day01.txt", "path to puzzle input file")
	flag.Parse()

	f, err := os.Open(*inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	switch *day {
	case 1:
		switch *part {
		case 1:
			ans, err := day01.Part1(lines)
			if err != nil {
				log.Fatalf("day1 part1 error: %v", err)
			}
			fmt.Println(ans)
		case 2:
			// placeholder for part 2
			fmt.Println("day1 part2 not implemented yet")
		default:
			log.Fatalf("unsupported part %d", *part)
		}
	default:
		log.Fatalf("unsupported day %d", *day)
	}
}
