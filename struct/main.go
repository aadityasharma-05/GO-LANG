package main

import "fmt"

func main() {
	fmt.Println("struct in golang")
	// no inheritance in golang ; no super or parent
	aditya:= User{"hitesh" , "aditya@go.com" , true , 19}
	fmt.Println(aditya)
	fmt.Printf(  "%+v",aditya)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


type order struct {
	Id string
	Amount float32 
	status string 
	
}