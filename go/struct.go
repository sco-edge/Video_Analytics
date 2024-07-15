package main

import "fmt"

type person struct { //struct 구조체는 이렇게 선언한다.
	name    string
	age     int
	favFood []string
}

func main_8() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}
