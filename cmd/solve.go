package main

import (
	"fmt"
	"github.com/somebadcode/adventofcode2019/internal/solver"
	"github.com/somebadcode/adventofcode2019/pkg/day1"
	"github.com/somebadcode/adventofcode2019/pkg/day2"
	"github.com/somebadcode/adventofcode2019/pkg/day3"
	"github.com/somebadcode/adventofcode2019/pkg/day4"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func solve(path string, config *viper.Viper, logger *log.Logger) error {
	wg := sync.WaitGroup{}

	solvers := []solver.Parts{
		day1.New(config.Sub("day1")),
		day2.New(config.Sub("day2")),
		day3.New(config.Sub("day3")),
		day4.New(config.Sub("day4")),
	}

	for i, s := range solvers {
		file, err := os.Open(filepath.Join(path, fmt.Sprintf("day%d.txt", i+1)))
		if err != nil {
			return err
		}

		wg.Add(1)
		go func(s solver.Parts, day int, f *os.File) {

			defer wg.Done()
			defer func() {
				if err := f.Close(); err != nil {
					fmt.Println(err)
				}
			}()
			p1 := s.PartOne(f)
			_, err := f.Seek(0, io.SeekStart)
			if err != nil {
				logger.Printf("Day %d\tError: %s\n", day, err)
				return
			}
			p2 := s.PartTwo(f)

			logger.Printf("Day %d\tPart 1: %s\n\tPart 2: %s\n", day, p1, p2)
		}(s, i+1, file)
	}

	wg.Wait()
	return nil
}
