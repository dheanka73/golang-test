package main

import "fmt"

func main() {

	var name string

	fmt.Print("What is your name ? : ")
	fmt.Scan(&name)

	if name == "batman" {
		fmt.Println("Your lying !")
	} else {
		fmt.Println("Oh, Hi ", name, " !")
	}

}
