package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		len   int
		line  string
		array []string
		check int
	)
	fmt.Scanln(&len)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
	array = strings.Fields(line)

	left := 0
	right := len - 1
	for {
		if left == right && len%2 != 0 {
			switch {
			case array[left] == "0":
				check++
			}
			break
		}

		if array[left] == "0" {
			check++
		}

		if array[right] == "0" {
			check++
		}

		left++
		right--

		if left > right && len%2 == 0 {
			break
		}

	}

	cnt := 0
	i := 0
	for {
		comp := array[i]
		switch {
		case comp == "0":
			tmp := array[i]
			array = append(array[:i], array[i+1:]...)
			array = append(array, tmp)
			cnt++
		case comp != "0":
			i++
		}
		if cnt == check {
			break
		}
	}

	result := strings.Join(array, " ")
	fmt.Println(result)
}
