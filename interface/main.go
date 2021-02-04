package main

import "fmt"

type I interface {
	i()
}

type A struct {
	AS string
}

func (a A) i() {

}

type B struct {
	BI int
}

func (b B) i() {

}

func main() {
	a := A{}
	b := B{}
	var ii I = a
	fmt.Println(ii.(B).BI)
	fmt.Println(b)
}
