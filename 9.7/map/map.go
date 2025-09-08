package main
import "fmt"

func main(){
	//(1)初始化
	stu1 :=map[string]string{"name":"geng","age":"20","gender":"mael"}
	fmt.Println(stu1)
	//(2)增删改查
	//1查看
	fmt.Println(stu1["name"])
	for i,v :=range stu1 {
		fmt.Println(i,v)
		
	}
	//2删除
	delete(stu1,"name")
	fmt.Println(stu1)
	//3添加和更新
	stu1["weeight"]="70kg"
	stu1["gender"]="男"
	fmt.Println(stu1)

	//进阶
	stu2 :=map[string]string{"name":"geng2","age":"20","gender":"mael"}
	stu3 :=map[string]string{"name":"geng3","age":"30","gender":"mael"}
	stu4 :=map[string]string{"name":"geng4","age":"40","gender":"mael"}
	var stus[]map[string]string
	stus = append(stus, stu2,stu3,stu4)
	fmt.Println(stus)
	fmt.Println(stus[1]["age"])

}
 