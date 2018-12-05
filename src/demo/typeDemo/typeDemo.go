package typeDemo

import (
	"fmt"
)

type Rect struct {
	width,length float64
}

type Phone struct {
	price int
	color string
}

type Iphone struct{
	Phone
	sys string
}

func (huawei *Phone) buy(){
	huawei.price += 1000
}

func Init(){
	var ip Phone
	ip.price = 6999
	ip.color = "red"
	ip.buy()
	fmt.Println(ip)
}

//定义在结构体上的方法
func (rectt *Rect) area() float64 {
	rectt.width *= 2
	rectt.length *= 2
    return rectt.width * rectt.length
}

func rectFunc()(sum float64){
	// var rect = new(Rect)
	var rect Rect
    rect.width = 100
	rect.length = 200
	fmt.Println(rect)
	sum = rect.area()
	fmt.Println(rect)
	return
}


