package main

import "fmt"

func main(){
    var p1,p2 *int
    i1, i2 := 123, 321
    p1,p2 = &i1, &i2
    fmt.Printf("%v %v\n", p1, p2)
    fmt.Println(p1 == p2)
    p2 = &i1
    fmt.Println(p1 == p2)
    var ap *[][]int
    //a := [][]int{{1},{2},{3}}
    //ap = &a
    ap = &[][]int{{1},{2},{3}}
    fmt.Printf("%v", ap)

    fmt.Println()
    fmt.Printf("addr p1 %p; addr &p1 %p; addr &i1 %p\n", p1, &p1, &i1)
}
