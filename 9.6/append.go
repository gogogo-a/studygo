package main

import "fmt"

func main() {
	var a =[]int{1,2,3}
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 2)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 3)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 4)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 5)
	fmt.Println(a, len(a), cap(a))

}