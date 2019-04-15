package main

import "fmt"

type aStruct struct {
    a int
    b string
    c *int
}


func main(){
    tmp := 651235
    var a = aStruct{a:1, b:"b", c:&tmp}
    fmt.Printf("%v\n", a)
}
