package main

import (
	"fmt"
	"github.com/somebadcode/adventofcode2019/pkg/day1"
	"github.com/somebadcode/adventofcode2019/pkg/day2"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func solve(path string, logger *log.Logger) error {
	wg := sync.WaitGroup{}

	functions := [][]func(io.ReadSeeker) string{
		{
			day1.PartOne,
			day1.PartTwo,
		},
		{
			day2.PartOne,
			day2.PartTwo,
		},
	}

	for i, parts := range functions {
		file, err := os.Open(filepath.Join(path, fmt.Sprintf("day%d.txt", i+1)))
		if err != nil {
			return err
		}

		wg.Add(1)
		go func(p1, p2 func(io.ReadSeeker) string, day int, f *os.File) {
			var resultOne, resultTwo string

			defer wg.Done()
			defer func() {
				if err := f.Close(); err != nil {
					fmt.Println(err)
				}
			}()

			resultOne = p1(f)

			_, err := f.Seek(0, io.SeekStart)
			if err != nil {
				panic(err)
			}

			resultTwo = p2(f)

			logger.Printf("Day %d\tPart 1: %s\n\tPart 2: %s\n", day, resultOne, resultTwo)
		}(parts[0], parts[1], i+1, file)

	}

	wg.Wait()
	return nil
}
