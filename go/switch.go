package main

import "fmt"

func canIDrink2(age int) bool {
	switch koreanAge := age + 2; { //if와 마찬가지로 내부에 선언가능, ; 을 빼먹으면 안 됨
	case koreanAge < 20:
		return false
	case koreanAge == 20:
		return true
	case koreanAge > 52:
		return false
	}
	return false
}

func main_4() {
	fmt.Println(canIDrink2(16))
}
