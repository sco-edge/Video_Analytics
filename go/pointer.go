package main

import "fmt"

func main_5() { //c언어와 마찬가지로 &을 쓰면 그 변수의 주소값이다. *는 주소값에 저장된 값이다.
	a := 2
	b := &a
	*b = 10
	fmt.Println(a, *b)
}
