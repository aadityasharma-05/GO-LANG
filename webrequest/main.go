package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://fakestoreapi.com/products/1"

func main() {
	fmt.Println("loc web request ")
	 res , err :=   http.Get(url)
	 
	 if err != nil {
		panic(err) 
	 } 

	 defer res.Body.Close() ; 

	 databytes , err :=  ioutil.ReadAll(res.Body) 
	 content := string(databytes)
 

	 fmt.Println(content)

} 