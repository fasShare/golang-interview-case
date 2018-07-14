package main

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

// 通过 go build -gcflags=-m closure_variable.go
// 查看x最终被转移到heap上，因此返回的func最终引用的只是x的地址

func main() {
	a, b := test(100)
	b()
	a()
}
