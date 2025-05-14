package main 

import "fmt"
// func main() {

// 	fmt.Println("msps in golang")

// 	lang := make(map[string]string) 

// 	lang["js"] = "javascript"
// 	lang["mb"] = "mongodb"
// 	lang["nj"] = "nodejs" 

// 	fmt.Println(lang) ;
	
// 	fmt.Println(lang["js"]) ;

// 	delete(lang , "mb")

// 	fmt.Println(lang) ;

 
// }

func main () {
	studentgrade := make(map[string]int) ;
	studentgrade["aditya"] = 122 ; 
	studentgrade["adi"] = 22 ; 
	studentgrade["riya"] = 92 ; 
	studentgrade["kabira"] = 67 ; 

	fmt.Println(studentgrade)

	// checking if a key exist 
	grades , exists := studentgrade["bhika"] 
	fmt.Println(grades)
	fmt.Println(exists)


}