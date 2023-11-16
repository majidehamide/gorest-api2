package main

import "fmt"

func Run() error {
	fmt.Println("Application run without error")
	return nil
}
func main() {
	fmt.Println("start the application")
	if err := Run(); err != nil {
		fmt.Println("Error when start application")
	}
}
