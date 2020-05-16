package main

import (
	"fmt"
	"reflect"
)

func main() {
	test(&Item{v: "123"})
	t2("12#")
}

func t2(x interface{}) {
	item := x.(string)
	fmt.Println(reflect.TypeOf(item))
}

type Item struct {
	v string
	i int
}

func test(x interface{}) {
	item := x.(*Item)
	fmt.Println(reflect.TypeOf(item))
}
