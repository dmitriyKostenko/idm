package main

import "fmt"

func main() {

	if err := someFunction(); err != nil {
		fmt.Println("Error occurred")

		return
	}
}

func someFunction() error {
	return fmt.Errorf("this is an error")
}
