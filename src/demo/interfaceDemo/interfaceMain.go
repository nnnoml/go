//结构体测试
package interfaceDemo

import(
	"fmt"
)

//接口
type Phone interface {
	sales() int
}

//结构体
type Nokia struct{
	price int
}
//结构体的方法
func (n Nokia) sales() int{
	return n.price
}

//第二个结构体
type Iphone struct{
	price int
}
//第二个结构体的方法
func (i Iphone) sales() int{
	return i.price
}
//第三个结构体 其中有接口的属性
type Person struct {
    name   string
    age    int
    phones []Phone //接口
}
//第三个结构体的累加方法
func (pes Person) total_cost() int {
    var sum = 0
    for _, phone := range pes.phones {
        sum += phone.sales()
    }
    return sum
}

func Init(){
	//定义结构体数组
    var bought_phones = [5]Phone{
        Nokia{price: 350},
        Iphone{price: 5000},
        Iphone{price: 3400},
        Nokia{price: 450},
        Iphone{price: 5000},
	}
	//定义第三个结构体 ，将上面的结构体数组切片传入，（引用传递），如果不用切片 就是值传递，那么在接收方的arr长度必须一致
    var person = Person{"Jemy",25, bought_phones[:]}
    fmt.Println(person.name)
    fmt.Println(person.age)
    fmt.Println(person.total_cost())
}