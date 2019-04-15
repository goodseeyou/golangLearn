package main

import (
    "fmt"
    "outer"
)

func main(){
    t := outer.T{B:2}
    fmt.Printf("%v\n",t)
}
