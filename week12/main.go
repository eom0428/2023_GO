package main

import "fmt"

func main() {

	s := []int{0, 0, 0, 0, 0}

	s[4] = 99
	s[2] = 91

	for _, value := range s {
		fmt.Println(value)
	}

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

	test := [3]string{"hamin", "bambi", "noha"} //배욜 리터럴을 이용해서 test 배열 생성
	testS := test[0:2]
	testS2 := test[1:]
	testS2[0] = "yejun"
	fmt.Println(testS2)
	fmt.Println(testS, len(testS))
}
