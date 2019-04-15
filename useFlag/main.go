package main

import (
    "fmt"
    "flag"
    "strings"
)


func main(){
    var t = flag.String("t", "", "string to change")
    var s = flag.String("s", " ", "string for seperator")
    flag.Parse()
    fmt.Printf("%v\n", flag.Args())
    fmt.Printf("%s\n", *t)
    stringFields := strings.Fields(*t)
    fmt.Println(strings.Join(stringFields, *s))
}
