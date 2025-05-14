package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "zpp , aranfr , jbsniusn"

	parts := strings.Split(data, ".")
	fmt.Println(parts)
}
