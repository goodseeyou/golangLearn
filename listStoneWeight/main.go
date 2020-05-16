package main
import "fmt"


func main() {
    fmt.Println(lastStoneWeight([]int{2,7,4,1,8,1}))
}


func lastStoneWeight(stones []int) int {
    return bruce(stones)
}

func bruce(stones []int) int {
    fmt.Println(stones)
    switch len(stones) {
        case 0:
            return 0
        case 1:
            return stones[0]
        case 2:
            return abs(stones[1]-stones[0])
    }

    if stones[len(stones)-1] < stones[len(stones)-2] {
        swap(stones, len(stones)-1, len(stones)-2)
    }

    for i := 0 ; i < len(stones) -2 ; i++{
        if stones[i] > stones[len(stones) - 1] {
            swap(stones, i, len(stones) - 1)
            swap(stones, i, len(stones) - 2)
        } else if stones[i] > stones[len(stones) - 2] {
            swap(stones, i, len(stones) - 2)
        }
    }

    v := abs(stones[len(stones) -1] - stones[len(stones) - 2])
    if v != 0 {
        return bruce(append(stones[:len(stones)-2], v))
    }

    return bruce(stones[:len(stones)-2])
}

func swap(s []int, i,j int) {
    s[i], s[j] = s[j], s[i]
}

func abs(v int) int {
    if v < 0 {
        return v
    }

    return v
}
