package main

import (
	foodchain "exercism/go/food-chain"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	// fmt.Println(input)
	num, _ := strconv.Atoi(input)
	// fmt.Println(num)
	fmt.Println(foodchain.Verse(num))
}
