package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://fakestoreapi.com/products/1"

func main() {
	fmt.Println("welcome to handling in URl")
	fmt.Println(myurl)

	result , _ := url.Parse(myurl)
// fmt.Println(result.Scheme)
// fmt.Println(result.Host)
// fmt.Println(result.Path)
// fmt.Println(result.Port())
fmt.Println(result.RawQuery)

qparam := result.Query()
fmt.Printf("the types of query params are : %T\n" , qparam)
	 fmt.Println(qparam["description"])


	  
} 