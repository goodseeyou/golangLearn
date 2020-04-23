package main
import (
    "fmt"
    "time"
)

func run(c chan int, n string) {
    for v := range c {
        fmt.Println(v, len(c), n)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    c := make(chan int, 15)
    go run(c, "A")
    go run(c, "B")
    for v := 1 ; v <= 30 ; v++ {
        c<-v
    }
    for {
        if len(c) == 0 {
            break
        }
    }
    close(c)
    fmt.Println("Finished")
}
