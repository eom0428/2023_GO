package main

import (
	"fmt"
)

func swap(n1 int, n2 int) {
	temp := n1
	n1 = n2
	n2 = temp
}

func main() {
	a := 10
	b := 20

	c := &a // var c *int = &a
	fmt.Printf("%T\n", c)

	swap(a, b)

	fmt.Println(a, b)
	fmt.Println(a, b)

}
