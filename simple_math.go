package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(x int64, y int64) int64 {
	return x + y
}

func minus(x int64, y int64) int64 {
	return x - y
}

func main() {
	a := (os.Args[1])
	b := (os.Args[2])

	c, _ := strconv.ParseInt(a, 10, 64)
	d, _ := strconv.ParseInt(b, 10, 64)
	resAdd := add(c, d)
	resAdd1 := c + d
	resMinus := minus(c, d)
	fmt.Println("addition: ", resAdd, resAdd1)
	fmt.Println("subraction: ", resMinus)
}
