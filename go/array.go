package main

import "fmt"

func main() {
	names := []string{"nico", "lynn", "dal"} //배열 선언할 때 크기를 앞에다가 적는다. 크기를 안 적으면 선언 시 알아서 크기를 할당한다.
	names = append(names, "hi")              //append는 해당 배열을 수정하는 것이 아니라 새로 만든다.
	fmt.Println(names)
}
