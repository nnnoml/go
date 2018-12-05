package base

import (
    "fmt"
)

func Init(){
	x := "10"
	y := "100"
	swap(&x,&y)
	fmt.Println(x,y)
}
//交叉赋值
func swap(x *string,y *string){
	*x,*y = *y,*x
}


//for循环测试
func ForTest()  {
	var a = [...]int{1,2,3,4,5}
	for _,item := range a{
		fmt.Println(item)
	}
}
//切片
func sliceTest(){
	x := make([]int,5,10)
	x[4]=4
	fmt.Println(x)
	x = append(x,6,7,8,9,9,9)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
//函数返回值
func modelRet(arr ...int)(x int,y string) {
	x = 0
	for _,item := range arr{
		x +=item
	}
	y = "this is model"
	return
}

//函数返回值是匿名函数
func createEvenGenerator()(func() uint) {
	// nextEven := createEvenGenerator()
    // fmt.Println(nextEven())
    // fmt.Println(nextEven())
	// fmt.Println(nextEven())

	i := uint(0)
    return func() (retVal uint) {
        retVal = i
        i += 2
        return
    }
}

//异常
func loop(){
	//本函数执行结束执行 无论正常结束还是异常终止
	defer func() {
		//终止后，获取真正的异常内容，或者获取慌张结束传递的内容
		fmt.Println(recover())
    }()
	fmt.Println("start")
	arr := []int{1,2,3,4}
	//如果有异常 直接结束
	arr[3]=5
	//如果没有异常 可以手动慌张结束
    panic("It starts to rain cats and dogs")
}

//指针
func zz(){
	var x int
	x = 123
	var x_p *int
	x_p = &x
	fmt.Println(x,x_p,*x_p,&x_p,*&x_p)
}