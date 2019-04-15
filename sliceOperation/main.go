package main
import "fmt"

func main(){
    var sliceInt []int = []int{1,2,3,4,5,6}
    partA := sliceInt[:4]
    partB := sliceInt[2:]
    fmt.Println(partA, partB)
    partA[3] = 999
    //partA[4] = 1000 //panic
    fmt.Println(partA, partB)
    fmt.Printf("cap of partA: %d\n", cap(partA))
    fmt.Println(partA[:5])

    var arrayInt [5]int = [...]int{1,2,3,4,5}
    intA := arrayInt[:4]
    intB := arrayInt[2:]
    fmt.Println(intA, intB)
    intA[3] = 999
    fmt.Println(intA, intB)

    copy(sliceInt[3:], sliceInt[4:])
    fmt.Println(sliceInt)
}
