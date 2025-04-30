package main 

import "fmt"
func main() {

	fmt.Println("msps in golang")

	lang := make(map[string]string) 

	lang["js"] = "javascript"
	lang["mb"] = "mongodb"
	lang["nj"] = "nodejs" 

	fmt.Println(lang) ;
	
	fmt.Println(lang["js"]) ;

	delete(lang , "mb")

	fmt.Println(lang) ;


}