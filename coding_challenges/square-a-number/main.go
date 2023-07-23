package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		number, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Please give a number !")
			return
		}
		result := number * number
		fmt.Println(result)
	}
}
