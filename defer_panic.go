package main

import "fmt"

// panic先于defer执行，但是panic的向上传递是需要在defer执行完之后。
// 一个函数最终只会向上传递一个panic，后发生的panic会覆盖之前的panic

func multi_panic_defer() {
	// panic one执行之后panic two也就不会被执行了
	// 由于panic的执行咸鱼defer，但是defer发生于panic向上传递之前，
	// 因此panic one在被向上传递之前defer中的panic defer发生了，
	// 并且覆盖了panic one，因此最终向上传递被捕获的只有panic defer

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("panic defer")
	}()

	panic("panic one")
	panic("panic two") // 第一个panic后的代码就不会被执行
}

func defer_panic() {
	defer func() {
		fmt.Println("A")
	}()
	defer func() {
		fmt.Println("B")
	}()
	defer func() {
		fmt.Println("C")
	}()

	panic("异常")
}

func main() {
	multi_panic_defer()
	fmt.Println("-------------------------")
	defer_panic()
}
