package main

import (
	"fmt"
	"log"
)

// 입력된 수의 소수 판정 프로그램 v0.6
func main() {
	var number int
	fmt.Print("정수 입력:")
	n, err := fmt.Scanln(&number)

	fmt.Println(n)

	if err != nil {
		log.Fatal(err)
	}

	isPrime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break // 첫 번째 약수가 발견되면 반복문 즉시 종료
		}
		//fmt.Print(i, " ")
	}

	if isPrime { // 비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}
