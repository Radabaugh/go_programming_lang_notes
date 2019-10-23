// Echo 2 prints its command-line argumetns.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	output, separation := "", ": "
	for i := 0; i < len(os.Args); i++ {
		output = strconv.Itoa(i) + separation + os.Args[i]
		fmt.Println(output)
	}
}
