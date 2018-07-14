package main

const cl = 100

var bl = 123

func main() {
	println(&bl, bl)
	// error 不能取常量地址，常量通常在编译期展开了
	println(&cl, cl)
}
