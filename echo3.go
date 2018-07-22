package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Index: %d Value: %s\n", i, os.Args[i])
	}
}
