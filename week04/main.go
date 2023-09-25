package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var a int    //declaration
	//a = 9        //assign value

	//var a int = 9

	//var a = 9
	a := 9                            //short form.declaration & assign value
	fmt.Println(a, reflect.TypeOf(a)) //declaration & assign value
	//b := 2.71
	var b float32 = 2.71
	c := 'Z'
	d, e := 4.10, 3.99
	f := "문자열"

	var g int
	var h float32
	var i bool
	var j string
	K := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할 수 있음"
	//koreanzombie := "정찬성"
	koreanZombie := "정찬성"
	fmt.Println(j, g, h, i, K, koreanZombie)

	fmt.Println(float32(a) * b)
	fmt.Println(a * int(b))     //여기서만 b가 int로 바뀌고
	fmt.Println(float32(a) > b) // 여기서는 b가 계속 float32임
	fmt.Println(a, reflect.TypeOf(a), b, reflect.TypeOf(b), c, reflect.TypeOf(c), d, reflect.TypeOf(d), e, reflect.TypeOf(e),
		f, reflect.TypeOf(f))

	//fmt.Println(reflect.TypeOf('Z'))
	//fmt.Println(reflect.TypeOf(99))
	//fmt.Println(reflect.TypeOf(2.7))
	//fmt.Println(reflect.TypeOf(false))
	//fmt.Println(reflect.TypeOf("Go!"))
	//fmt.Println('A', 'a', '0', '김', '인', '하') //65, 97, 48, 44608, 51064, 54616
	//fmt.Println(math.Ceil(3.17))
	//fmt.Println(math.Floor("3.14"))
	//fmt.Println(strings.Title("go!"))
	//fmt.Println(strings.Title("open source programming~\n\"go\"!"))
}
