package main

import "fmt"

func process(kor int, eng int, math int) (int, int) {
	totalScore := kor + eng + math
	average := totalScore / 3

	return totalScore, average
	//fmt.Printf("%s님의 총점을 %d점, 평균은 %d점 입니다.\n", name, totalScore, average)
}

func main() {
	t, a := process(95, 88, 70)
	fmt.Printf("%s님의 총점을 %d점, 평균은 %d점 입니다.\n", "남예준", t, a)

	t, a = process(95, 86, 88)
	fmt.Printf("%s님의 총점을 %d점, 평균은 %d점 입니다.\n", "남예준", t, a)

}
