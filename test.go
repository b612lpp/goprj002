package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}

	for c, v := range a {
		fmt.Println(v, c)
	}
}
