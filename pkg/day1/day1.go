package day1

import (
	"bufio"
	"io"
	"strconv"
)

func PartOne(r io.Reader) string {
	var fuelQuantity int64

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err.Error()
		}

		mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return err.Error()
		}

		fuelQuantity += (mass / 3) - 2
	}

	return strconv.FormatInt(fuelQuantity, 10)
}

func PartTwo(r io.Reader) string {
	var fuelQuantity int64

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var totalFuelRequirement int64

		if err := scanner.Err(); err != nil {
			return err.Error()
		}

		mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return err.Error()
		}

		baseFuelRequirement := (mass / 3) - 2
		q := baseFuelRequirement
		for {
			totalFuelRequirement += q
			q = (q / 3) - 2
			if q <= 0 {
				break
			}
		}

		fuelQuantity += totalFuelRequirement
	}

	return strconv.FormatInt(fuelQuantity, 10)
}
