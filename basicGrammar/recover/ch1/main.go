package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

/*
用法:使用recover()捕捉panic异常的时候,则需要defer来读取一个匿名函数
理论:如果不进行recover,便会导致整个程序挂掉,
Recover()用法是:将Recover()写在defer中,并且在可能发生panic的地方之前,先调用此defer的东西（让系统方法域结束时,有代码要执行.）
当程序遇到panic的时候(当然,也可以正常的调用出现的异常情况)系统将跳过后面的代码,进入defer,如果defer函数中recover(),则返回捕获到的panic的值。
*/

func test2() {
	fmt.Println("正常程序")
}
func test1(x int) {
	defer func() {
		if err := recover(); err != nil {
			//错误输出
			fmt.Println("程序异常", err)
			//将Stack返回信息打印到标准错误输出
			debug.PrintStack()
		}
	}()
	var a [10]int
	a[x] = 1222
	log.Println(a)
}
func main() {
	test1(20)
	//即使上面函数panic了,但是recover机制恢复了程序正常执行
	test2()
}
