package main

type defintion_int int // 新定义一个类型，赋值需要强转
type alias_int = int   // 定义类型别名，是同一个类型

func (self *defintion_int) add(dela defintion_int) defintion_int {
	*self += dela
	return *self
}

// golang原生类型以及其别名上定义方法，
//func (self *alias_int) add(dela int) int {
//	*self += dela
//	return *self
//}

func main() {
	var tmp int = 99
	var as alias_int = tmp
	var di defintion_int = defintion_int(tmp) // error
	println(as, di)

	println(di.add(1))
}
