package main

import (
	"fmt"
	"week10/src/greeting"
	"week10/src/mymath"
)

func main() {
	fmt.Println(mymath.MyPower(2, 9))
	fmt.Println(mymath.MyAbs(mymath.MyAbs(-99)))
	greeting.Hello()
	greeting.Hi()
}
