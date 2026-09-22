package main

import (
	"fmt"
	"os"
)

func main() {
	err := readInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error", err)
		return
	}

}
