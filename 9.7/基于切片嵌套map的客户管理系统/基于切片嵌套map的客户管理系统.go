package main

import "fmt"

func main() {
	userlist := []map[string]interface{}{}

	for {
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
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("查看用户")
			fmt.Println(userlist)
		case 2:
			user := make(map[string]interface{})
			var name string
			var age string
			var gender string
			fmt.Print("添加用户,请输入用户名称")
			fmt.Scanln(&name) // 使用 &name 传递变量地址，并用 Scanln
			fmt.Print("添加用户,请输入用户年龄")
			fmt.Scanln(&age) // 使用 &name 传递变量地址，并用 Scanln
			fmt.Print("添加用户,请输入用户性别")
			fmt.Scanln(&gender) // 使用 &name 传递变量地址，并用 Scanln
			user["name"] = name
			user["age"] = age
			user["gender"] = gender
			userlist = append(userlist, user)
			fmt.Printf("用户%s已添加成功\n", name)
		case 3:
			fmt.Print("删除用户,请输入用户名称")
			var name string
			fmt.Scanln(&name) // 使用 &name 传递变量地址，并用 Scanln
			var n int
			for i, v := range userlist {
				if v["name"] == name {
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
			fmt.Print("编辑用户,请输入用户名称")
			var name string
			fmt.Scanln(&name) // 使用 &name 传递变量地址，并用 Scanln

			var n int
			for i := range userlist {
				if userlist[i]["name"] == name {
					n = 1
					fmt.Printf("这是用户信息请核对%s %s %s\n", userlist[i]["name"], userlist[i]["age"], userlist[i]["gender"])
					for true {
						fmt.Println(`
						请选择需要修改的项
						1.name
						2.age
						3.gender
						4.退出
						`,
						)
						var s int
						fmt.Scanln(&s)
						switch s {
						case 1:
							fmt.Println("请输入修改后的姓名：")
							var name string
							fmt.Scanln(&name)
							userlist[i]["name"] = name
						case 2:
							fmt.Println("请输入修改后的年龄：")
							var age string
							fmt.Scanln(&age)
							userlist[i]["age"] = age
						case 3:
							fmt.Println("请输入修改后的性别：")
							var gender string
							fmt.Scanln(&gender)
							userlist[i]["gender"] = gender
						case 4:
							goto EditLoopEnd
						default:
							fmt.Println("无效的选择")
						}
					}
				
				}
			}
		EditLoopEnd:
			if n == 1 {
				fmt.Printf("用户%s已修改成功", name)
			} else {
				fmt.Printf("未找到用户%s", name)
			}

		case 5:
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("无效选择")

		}

	}
}
