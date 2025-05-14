package main

import "fmt"

//  by value
// func changeNumber(num int ) {
// 	num = 5
// 	fmt.Println("in changeNum" , *num )
// }

// referace by

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNUm", *num)
}

func main() {

	num := 1
	changeNum(&num)
	fmt.Println("the value is ", num)
}
