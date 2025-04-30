package main

import (
	"fmt"
	"sort"
)

func main() {
	number := []int{1, 2, 3, 4, 5}
	number = append(number , 3,3,43,2,2,32,3,2,2,1)
	fmt.Println("number : ", number)
	fmt.Printf("number  is type of data type  : %T\n ", number)

	// make function  

	highsor := make([]int , 4 ) 

	highsor[0] = 11 
	highsor[1] = 114 
	highsor[2] = 112

	highsor = append( highsor , 23,2,2) 

	fmt.Println(highsor)

	sort.Ints(highsor)
	fmt.Println(highsor)

	var tech = [] string{"react" , "node" , "mongodb" , "express" , "nosql"}
     var index int = 2 

	 tech = append(tech[:index] , tech[index+1:]...)  
	 fmt.Println(tech)
}