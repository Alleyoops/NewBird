package main

import "fmt"

// 人类结构体
type person1 struct {
	name   string // 姓名
	age    int    // 年龄
	gender string // 性别
}

// 复读机接口
type repeater1 interface {
	repeat(string)
}
//鸽子接口
type dove1 interface{
	gugugu()
}
//柠檬精接口
type lemoner interface{
	suan()
}
//真相怪接口
type zhenxiangmonster interface{
	xiang()
}

// 方法
func (p person1) repeat(word string) {
	fmt.Println(word)
}
func (p person1) gugugu(){
	fmt.Println("鸽了")
}
func (p person1) suan(){
	fmt.Println("酸")
}
func (p person1) xiang(){
	fmt.Println("真香")
}

func main() {
    p:=person1{}

	var	r repeater1
	var d dove1
	var	l lemoner
	var	z zhenxiangmonster

    r=p
    r.repeat("复读")
    d=p
    d.gugugu()
    l=p
    l.suan()
    z=p
    z.xiang()

}