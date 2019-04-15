package main
import "fmt"


func MultipleNamedResult(x,y int) (rx,ry int, err error) {
    if x < 0 || y < 0 {
        err = fmt.Errorf("x, y should > 0")
        return
    }
    /*
    // delcare function in a function is NOT allowed
    func deferAsign(x,y int) {
        rx, ry = x+y, x-y
    }*/

    // But anonymous function is allowed
    defer func(){rx,ry = x+y, x-y}()
    return 100,200, nil
}


func main(){
    x, y, _ := MultipleNamedResult(3,-1)
    fmt.Println(x,y)
}
