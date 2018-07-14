package main

func main() {
	println(DeferFunc1(1)) // 4
	println(DeferFunc2(1)) // 1
	println(DeferFunc3(1)) // 3
}

// defer、return、返回值三者的执行逻辑应该是：
// return最先执行，return负责将结果写入返回值中；
// 接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
