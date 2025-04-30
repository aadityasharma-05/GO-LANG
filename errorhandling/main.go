package main

import "fmt"

// func divide(a, b float64) (float64, error) {

// 	if b == 0 {
// 		return 0, fmt.Errorf("b cant be 0") ;
// 	}

// 	return a / b, nil
// }

// return string 

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "b cant be 0" ;
	}

	return a / b," nil"
}

func main() {
	fmt.Println("error handling ")
	ans, _ := divide(10, 5)
	fmt.Println(ans)
}
