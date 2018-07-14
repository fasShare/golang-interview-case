package main

import "fmt"

var (
	// size := 1024 // error
	size     = 1024
	max_size = size * 2
)

// := 变量的简短模式
// (1)定义变量的同时进行初始化
// (2)只能用于函数内部
// (3)不能提供数据类型

func main() {
	fmt.Println(size, max_size)
}
