package main
import "fmt"

func main() {
a:="bxj##tw"
b:="bxj###tw"
    fmt.Println(backspaceCompare(a,b))
}

func backspaceCompare(S string, T string) bool {
    s := genSlice(S)
    t := genSlice(T)

    fmt.Println(s)
    fmt.Println(t)

    if len(s) != len(t) {
        return false
    }

    for i, v := range s {
        if v != t[i] {
            return false
        }
    }

    return true
}

func genSlice(s string) []byte {
    skipCount := 0
    slice := []byte{}
    for i:=len(s)-1 ; i >= 0 ; i-- {
        if s[i]==byte("#"[0]) {
            skipCount++
        } else if skipCount > 0 {
            skipCount--
        } else {
            slice = append(slice, s[i])
        }
    }
    return slice
}
