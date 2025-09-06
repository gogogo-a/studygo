package main

import "fmt"

func main() {
	
	s := []int{1, 2, 3, 4,5}
	fmt.Println(s)
	//(1)移动到第一个位置
	// s = append([]int{100}, s...)
	// fmt.Println(s)
	// n := []int{11, 22, 33}
	// s=append(n,s... )
	// fmt.Println(s)
	//(2)任意位置添加
	// s=append(s[:2],append([]int{0,99},s[2:]...)...)
	// fmt.Println(s)
	//（3）删除某索引
	// s=append(s[:2],s[3:]... )
	// fmt.Println(s)
	//按值删除
	// 方式一
	// var s2 = []int{}
	// for _, v := range s {
	// 	if v != 1 {
	// 		s = append(s2, v)
	// 	}
	// }
	// fmt.Println(s)
	// 方式二
	for i, v := range s {
		if v == 1 {
			s = append(s[:i], s[i+1:]...)
			break // 删除后需要退出循环，否则索引会错乱
		}
	}
	fmt.Println(s)
		
	}

