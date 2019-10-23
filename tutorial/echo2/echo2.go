// Echo 2 prints its command-line argumetns.
package main

import (
	"fmt"
	"os"
)

func main() {
	output, separation := "", ""
	for _, arg := range os.Args[1:] {
		output += separation + arg
		separation = " "
	}
	fmt.Println(output)
}
