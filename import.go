package main

import (
	"bufio"
	"os"
	"strconv"
)

func readInput() error {
	var numbers []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		numbers = append(numbers, number)

	}

	return scanner.Err()
}
