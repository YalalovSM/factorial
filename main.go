package main

import (
	"log"

	"github.com/yalalovsm/factorial/server"
)

func main() {
	s := server.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
