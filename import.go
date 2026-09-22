package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func readInput() error {
	var numbers []float64

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.Atoi(line)

		if err != nil {
			return err
		}

		numbers = append(numbers, float64(number))

		prediction := float64(number)

		if len(numbers) >= 2 {
			a, b := linearRegression(numbers)

			nextX := float64(len(numbers))

			prediction = a*nextX + b

		}
		lower := int(math.Floor(prediction - 10))

		upper := int(math.Ceil(prediction + 10))

		fmt.Println(lower, upper)

	}

	return scanner.Err()
}
