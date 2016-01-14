package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	start := time.Now()

	fmt.Println(os.Args[1:])

//func make([]T, len, cap) []T

	fmt.Printf("\nThis took %v\n", time.Since(start))
}