package main

import (
	"fmt"
	"runtime"
	"sync"
)

func test_one() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	// 最终每个函数的输出内容一样，都是i=2时的变量地址和值
	return funs
}

func test_two() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	// x每次循环都会不同，所以返回数组中的每一个函数输出都是不一样的
	return funs
}

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	funs := test_one()
	for _, f := range funs {
		f()
	}

	funs = test_two()
	for _, f := range funs {
		f()
	}
}
