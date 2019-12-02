package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	flag.Parse()
	if flag.NArg() != 1 {
		logger.Fatalln("Please specify the directory where the input files (day1.txt, day2.txt etc)")
	}

	err := solve(flag.Arg(0), logger)
	if err != nil {
		os.Exit(1)
	}
}
