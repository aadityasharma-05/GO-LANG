package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("jns")
	// }

	nnumber := []int{1,2,3,4,5,6,7,9} 

	 for index , value := range nnumber {
		fmt.Println("index %d , value %d\n" , index , value ) 
	 }
}