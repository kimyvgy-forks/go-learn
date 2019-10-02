package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name := "Kimyvgy"
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Hello, what is your name?")
	fmt.Println("Hi there! My name is " + name + "!")

	fmt.Println("Where are you?")
	fmt.Println("I am at: ", cwd)
}
