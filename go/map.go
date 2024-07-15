package main

import "fmt"

func main() {
	nico := map[string]string{"name": "nico", "age": "12"} //[]안에 key, 그 옆에 value를 적는다.
	for _, value := range nico {
		fmt.Println(value)
	}
}
