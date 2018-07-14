package main

// goto不能跳转到其他函数或者内层代码

func main() {
	for i := 0; i < 10; i++ {
	loop:
		println(i)
	}
	goto loop
}
