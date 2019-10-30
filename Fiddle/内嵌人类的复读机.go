package main

import "fmt"

// person 人类
type person struct {
	name   string // 姓名
	age    int    // 年龄
	gender string // 性别
}

// repeater 复读机
type repeater第二版 struct {
	person
	repeatNum int // 复读次数
}

// 复读的函数即方法
func (r *repeater第二版) repeatagain(word string) {
	fmt.Println(word)
	r.repeatNum++
}

func main(){
	r:=repeater第二版{
		person:person{
			name: "复读机",
			age: 1998,
			gender: "other",
		},
		repeatNum: 99,
	}
	fmt.Println(r.name)
	r.repeatagain("全新版本复读机")
}
