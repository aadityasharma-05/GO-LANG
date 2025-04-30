package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("dice roll game")

	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6)
	fmt.Println("value of dice is ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("salman")
	case 2:
		fmt.Println("kabir ")

	case 3:
		fmt.Println("aditya ")
	case 4:
		fmt.Println("yuvisharma")

		default : 
		fmt.Println("no thanks ")
	}

	


}
