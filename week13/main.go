package main

import "fmt"

func main() {
	var a []string
	var b []bool
	a = make([]string, 4, 5)

	fmt.Println(b, len(b), cap(b))
	fmt.Printf("%#v %#v\n", a, b)
	fmt.Println(a, len(a), cap(a))
}
