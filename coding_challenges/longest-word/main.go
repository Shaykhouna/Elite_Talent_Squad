package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		list := strings.Split(os.Args[1], " ")
		if len(list) == 1 && os.Args[1] != "" {
			fmt.Println(os.Args[1])
		} else {
			count, longest := 0, "string"
			// grouped := make(map[string]int)
			for _, word := range list {
				size := 0
				for size < len(word) {
					size++
				}
				if count < size {
					count = size
					longest = word
				}
				size = 0
			}

			fmt.Println(longest)
		}
	}
}
