package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("get working directory failed: %v", err)
	}

	fmt.Println("current working directory: ", dir)
}
