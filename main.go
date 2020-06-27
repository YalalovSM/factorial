package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/yalalovsm/factorial/calculations"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: main.go <number>")
		return
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Please pass a number as an argument")
	}

	// f := calculations.FactorialNaive(n)
	f := calculations.FactorialTree(n)
	fmt.Println(f)
}
