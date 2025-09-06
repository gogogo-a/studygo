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
	// var s= []int{1,2,3}
	// s2 :=append(s,4)
	// fmt.Println(s2)
	// s[0]=100
	// fmt.Println(s2)
	var s=make([]int ,3,6)
	s2:=append(s,4)
	fmt.Println(s2)
	s[0]=100
	fmt.Println(s2)
	// 容量不够是新的数组不共用同一块空间，容量够就是共用一块空间
	a=append(a, []int{11,22,33}...)
	fmt.Print(a)

	L := make([]int, 5, 10)
	v1 := append(L, 1)
	fmt.Println(v1)
	fmt.Printf("%p\n", &v1)
	v2 := append(L, 2)
	fmt.Println(v2)
	fmt.Printf("%p\n", &v2)
	fmt.Println(v1)

	//经典面试题
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2] // [10, 20]
	s2 = s1 // [10, 20]
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 1000
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(arr)
}