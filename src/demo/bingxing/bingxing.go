package bingxing

import (
    "fmt"
    "math/rand"
	"time"    
	"strconv"
)

//协程 睡一会 并行执行
func list_elem(n int, tag string) {
    for i := 0; i < n; i++ {
        fmt.Println(tag, i)
        tick := time.Duration(rand.Intn(100))
        time.Sleep(time.Millisecond * tick)
    }
}
func Init_elem() {
    go list_elem(10, "go_a")
    go list_elem(20, "go_b")
    var input string
    fmt.Scanln(&input)
}


//投篮模型
//两分球，投三次停止
func fixed_shooting(msg_chan chan string) {
    var times = 3
    var t = 1
    for {
        if t <= times {
            msg_chan <- "two shooting"
        }
        t++
        time.Sleep(time.Second * 6)
    }
}
//三分球，投五次停止
func three_point_shooting(msg_chan chan string) {
	var times = 5
    var t = 1
    for {
        if t <= times {
            msg_chan <- "three point shooting"
        }
        t++
        time.Sleep(time.Second * 1)
    }
}
//四分球 channel缓冲
func four_point_shooting(msg_chan chan string){
	var group = 1
    for {
        for i := 1; i <= 10; i++ {
            msg_chan <- strconv.Itoa(group) + ":" + strconv.Itoa(i)
        }
        group++
        time.Sleep(time.Second * 10)
    }
}
//投球累加
func count(msg_chan chan string) {
    for {
        msg := <-msg_chan
        fmt.Println(msg)
        //time.Sleep(time.Second * 1)
    }
}
func Init() {
	//两分球的channel
	c := make(chan string)
	//三分球的channel
	c3 := make(chan string)
	//四分球的channel 有长度的channel 异步 channel缓冲区满了 或者调用完毕 执行读取操作
	c4 := make(chan string,20)
	//协程 启动两分投篮
	go fixed_shooting(c)
	//协程 启动三分投篮
	go three_point_shooting(c3)
	//协程 启动四分投篮
	go four_point_shooting(c4)
	go count(c4)
    go func() {
        for {
			//监听channel是否有值入栈
            select {
				//出栈
            case msg1 := <-c:
                fmt.Println(msg1)
            case msg2 := <-c3:
				fmt.Println(msg2)
				//监听超时提示
			case <-time.After(time.Second * 5):
				fmt.Println("timeout, check again...")
			}
        }
    }()
    var input string
    fmt.Scanln(&input)
}