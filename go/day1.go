package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) { //함수 옆에 있는 parameter는 입력값, 그 옆은 출력값이다.
	defer fmt.Println("i'm done") //defer는 함수가 종료되고 실행된다.
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	tlen, up := lenAndUpper("nico")
	fmt.Println(tlen, up)
}
