package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		len     int
		line    string
		arr     []string
		element string
		result  string
	)

	fmt.Scanln(&len)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
	fmt.Scanln(&element)
	arr = strings.Fields(line)

	i := 0
	check := 0
	for {
		comp := arr[i]
		switch {
		case comp == element:
			for j := i; j < len-1; j++ {
				arr[j] = arr[j+1]
			}
			check++
		case comp != element:
			i++
		}
		if i == (len-check-1) && arr[i] == element {
			check++
			break
		}
		if i == (len - check - 1) {
			break
		}
	}

	for i := 0; i <= (len - check - 1); i++ {
		switch {
		case i == (len - check - 1):
			result = result + arr[i]
		case i != (len - check - 1):
			result = result + arr[i] + " "
		}
	}

	fmt.Println(result)
}
