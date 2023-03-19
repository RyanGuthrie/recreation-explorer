package main

import (
	"log"
	"server/domain"
	"server/prompt"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var cursorPos int = 0
	for {
		state, err := domain.GetState(&cursorPos)

		if err == prompt.ExitError {
			log.Println("Exiting, goodbye!")
			return
		} else if err != nil {
			log.Fatalln(err)
		}

		if err = state.Explore(); err != nil {
			log.Fatalln(err)
		}

	}
}
