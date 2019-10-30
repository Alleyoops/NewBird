package main
import "fmt"
func main(){
	var s string
	fmt.Println("输入一个字符串儿：")
	fmt.Scan(&s)
	new:=change(s)
	fmt.Println("逆序输出你的字符串儿：")
	fmt.Println(new)

}
//定义逆序输出函数
func change(str string) string {
	arr:=[]rune(str)
	long:=len([]rune(str))
	for r,l:=0,long-1 ; r<l ; r,l=r+1,l-1 {
		arr[r],arr[l]=arr[l],arr[r]
	}

	return string(arr)
}
