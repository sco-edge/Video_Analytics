package main

import "fmt"

func superAdd_1(numbers ...int) int {
	for idx, number := range numbers {
		fmt.Println(idx, number)
	}
	return 1
}
func superAdd_2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main_2() {
	superAdd_1(1, 2, 3, 4, 5, 6)
	superAdd_2(1, 2, 3, 4, 5, 6)
	res := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(res)
}
