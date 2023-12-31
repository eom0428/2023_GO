package main

import "fmt"

func main() {
	//var primes [3]int
	//primes[0] = 2
	//primes[1] = 3
	//primes[2] = 5
	//fmt.Println(primes, primes[1])

	//var primes [3]int = [3]int{2,3,5} //배열 리터럴로 초기화
	//fmt.Println(primes, primes[1])

	primes := [3]int{2, 3, 5} //단축연산자.
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) //boolean 타입의 제로값 출력

	i := 0
	//for i < 4
	for i < len(primes) {
		fmt.Println(primes[i])
		i++
	}

	//for j := 0; j < len(primes); j++ {
	//	fmt.Println(primes[j])
	//}

	for _, prime := range primes {
		fmt.Println(prime)
	}

	fmt.Printf("%#v\n", test)
	fmt.Println(test)
	fmt.Printf("%#v\n", primes)
	fmt.Println(primes)
}
