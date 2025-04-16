package main

import "fmt"
import "time"

func main() {

	var name string
	var number int

	fmt.Print("What is your name ? : ")
	fmt.Scan(&name)

	if name == "batman" {
		fmt.Println("Your lying !")
	} else {
		fmt.Println("Oh, Hi ", name, " !")

		fmt.Print("Type any number, ", name, "!: ")
		fmt.Scan(&number)

		fmt.Println("Um......")
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("You know......")
			time.Sleep(1 * time.Second)

			fmt.Print()
			for i := 0; i < number; i++ {
				fmt.Println("I love you !")
			}
		}()
		time.Sleep(time.Second * 3)
	}

}
