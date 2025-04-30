package main
import "fmt"


func simpleheii() {
	fmt.Println("heelo")
}

func add( z , x int )(int) {
	return z+x ;
}

func aadd( z , x int )( result int) {
	result =  z+x ;
	return
}
func main() {
	fmt.Println("hello")
	simpleheii()
     sum := add(1,2 ) ;
	 fmt.Println(sum)
	 summ := aadd(1,2 ) ;
	 fmt.Println(summ)
}
