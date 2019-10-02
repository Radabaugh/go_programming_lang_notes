package main

import (
	"fmt"
	"os"
)

func main() {
	var output, separation string
	const usage = "Usage: echos back command-line arguments."

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			output += separation + os.Args[i]
			separation = " "
		}
		fmt.Println(output)
	} else {
		fmt.Println(usage)
	}
}
