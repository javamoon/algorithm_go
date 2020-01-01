package main

import (
	"fmt"
)

type Node0 struct {
	num int
}

type Node struct {
	val int
	next Node
}

func main() {

	nums := []int{9, 1, 2, 3, 5, 5, 5, 5, 5, 4, 5, 8, 2, 9, 12, 13, 15, 2, 16, 16, 16, 16, 16, 16, 16, 2, 15, 15, 15}

	i := Duplication(nums)

	if i > -1 {
		logx("length", len(nums))
		logx("i", i)
		logx("v", nums[i])
	} else {
		log("error")
	}

	fmt.Println("ok")
}

func logx(str string, n int)  {
	fmt.Printf("%s %d\n", str, n)
}

func logn(n int) {
	fmt.Printf("-> %d\n", n)
}

func log(x interface{}) {
	fmt.Println(x)
}
