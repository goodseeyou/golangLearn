package main
import "fmt"

func main(){
    fmt.Printf("%T %s\n", 'x', "x"[0])
    str := "0123456789"
    str = "中文"
    fmt.Println(len(str))
    fmt.Printf("%v", str)
    fmt.Printf(`backquoted %% \t 字串\n`)
}
