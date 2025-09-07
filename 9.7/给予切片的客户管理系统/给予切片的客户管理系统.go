package main

import "fmt"

func main() {
	userlist := []string{}
	for true {
		fmt.Println(
			`
			1.查看用户
			2.添加用户
			3.删除用户
			4.编辑用户
			5.退出系统
			`,
		)
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("查看用户")
			fmt.Println(userlist)
		case 2:
			fmt.Print("添加用户,请输入用户名称")
			var name string
			fmt.Scanln(&name) // 使用 &name 传递变量地址，并用 Scanln
			userlist = append(userlist, name)
			fmt.Printf("用户%s已添加成功\n", name)
		case 3:
			fmt.Print("删除用户,请输入用户名称")
			var name string
			fmt.Scanln(&name) // 使用 &name 传递变量地址，并用 Scanln
			var n int
			for i, v := range userlist {
				if v == name {
					userlist = append(userlist[:i], userlist[i+1:]...)
					n = 1
					break
				}
			}
			if n == 1 {
				fmt.Printf("用户%s已删除成功", name)
			} else {
				fmt.Printf("未找到用户%s", name)
			}

		case 4:
			fmt.Println("编辑用户")
		case 5:
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("无效选择")

		}

	}
}
