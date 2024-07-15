package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 { //if 내에 변수 선언 가능. 이 변수는 if 문을 처리할 때만 사용하기 위해 선언한 것임을 알 수 있게 함
		return false
	}
	return true
}

func main_3() {
	fmt.Println(canIDrink(16))
}
