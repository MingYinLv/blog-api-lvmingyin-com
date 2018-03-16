package main

import (
	"fmt"
)

func main() {
	for i := range []int{1,2,3}{
		fmt.Println(i)
	}
}
