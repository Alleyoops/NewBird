package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {

	// 图片大小
	const size,sizex = 600,500
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, sizex, size))

	// 遍历每个像素
	for x := 0; x < sizex; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	// 从0到最大像素生成x坐标
	for x := 0; x <sizex; x++{
		//上面1/4部分
		// 让自变量s的值的范围在0~2Pi之间
		s := float64(x)/sizex *2* math.Pi
		// sin的图像取绝对值并经过翻转平移
		y := size/4-math.Abs(math.Sin(s))*size/4
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})

		//下面3/4部分
		n := float64(x)/sizex * math.Pi
		m:= size/4+math.Sin(n)*size/4*3
		pic.SetGray(x, int(m), color.Gray{0})


		// 创建文件
		file, err := os.Create("sin.png")

		if err != nil {
			log.Fatal(err)

		}
		// 使用png格式将数据写入文件
		png.Encode(file, pic) //将image信息写入文件中

		// 关闭文件
		file.Close()

	}
	fmt.Printf("程序生成的PNG文件在该go文件所在文件夹，请在return 0后打开查看\n"+
		"P.S\n"+"借鉴网上输出正弦函数的程序，自己用了两个正弦函数拼接而成\n" +
		"上面的sin取了绝对值，两个函数衔接的不是太美妙，勉强有大概的一个形状")
}
