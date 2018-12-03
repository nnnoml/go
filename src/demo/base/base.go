package base

import (
    "fmt"
)

func Init(){
	sliceTest()
}

//for循环测试
func ForTest()  {
	var a = [...]int{1,2,3,4,5}
	for _,item := range a{
		fmt.Println(item)
	}
}

func sliceTest(){
	x := make([]int,5,10)
	x[4]=4
	fmt.Println(x)
	x = append(x,6,7,8,9,9,9)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}