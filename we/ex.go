package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	prime := true

	if n < 2 {
		return false, fmt.Errorf("%d 는 소수가 아닙니다.", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}
	return prime, nil
}

func main() {
	var a, b int

	fmt.Print("정수 입력: ")
	_, err := fmt.Scanln(&a, &b)

	if err != nil {
		log.Fatal(err)
	}

	for i := a; i <= b; i++ {

		p, err := isPrime(i)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			fmt.Print(i, " ")
		}
	}

}
