# Advent of Code 2025

Go solutions for [Advent of Code 2025](https://adventofcode.com/2025), one folder per day.

## Layout

- `cmd/aoc2025`: main runner to execute a specific day/part.
- `internal/day01`: solution code and tests for Day 1.
- `input/`: puzzle input files (not committed here by default).

## Running

Run Day 1, Part 1 with your puzzle input in `input/day01.txt`:

```bash
cd /Users/moritzjo/GolandProjects/AdventOfCode2025

go test ./...

go run ./cmd/aoc2025 -day=1 -part=1 -input=input/day01.txt
```

Add future days under `internal/dayXX` and extend the `main.go` switch as you go.

