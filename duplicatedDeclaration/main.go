package main
import "fmt"

var aInt int = 456

func main(){
    fmt.Printf("%p %v\n",&aInt, aInt)
    aInt := 123
    fmt.Printf("%p %v\n",&aInt, aInt)
}


