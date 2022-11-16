package main

import "fmt"

// Run - is going to be responsible for
// the instantation and startup of our
// go application

func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
