package main
import "fmt"

// dove 鸽子
type dove struct {
	name string
	age int
	gender string
	doveNum int    // 鸽的次数
}

// repeater 复读机
type repeater struct {
	name1 string
	age1 int
	gender1 string
	repeatNum int    // 复读次数
}

// 复读的函数即方法
func (r *repeater) repeat(word string) {
	fmt.Println(word)
	r.repeatNum++
}

//鸽的函数即方法
func (d *dove) gugugu() {
	fmt.Println(d.name,"又鸽了")
	d.doveNum++
}

func main (){

//鸽和复读结构体的实例化（声明）
	d:= dove{
	name:    "臭鸽子",
	age:     21,
	gender:  "male",
	doveNum:999,// 如果分多行声明结构体，每个参数后面都必须加逗号
}
	r:=repeater{
		name1: "复读机",
		age1: 1998,
		gender1: "female",
		repeatNum: 99,
	}

//开始鸽和复读
	d.gugugu()
	r.repeat("我又复读了耶")
	fmt.Printf("鸽子鸽了第%d次，复读机复读了第%d次",d.doveNum,r.repeatNum)
	}


