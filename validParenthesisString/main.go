package main
import "fmt"

func checkValidString(s string) bool {
    stack := make([]string, 0)

    for i:=0; i<len(s) ; i++ {
        if string(s[i]) == "(" || string(s[i]) == "*" {
            stack = append(stack, string(s[i]))
        } else if string(s[i]) == ")" {
            // find the last "("
            j:=len(stack) - 1
            for ; j>=0 ; j-- {
                if stack[j] == "(" {
                    break
                }
            }

            if j < 0 {
                if len(stack) < 1 {
                    fmt.Println("No ( 1st")
                    return false
                } else {
                    // remove the last star
                    stack = stack[:len(stack)-1]
                }
            } else {
                // remove the last "("
                for k:=j; k<len(stack) - 1 ; k++ {
                    stack[k] = stack[k+1]
                }
                stack = stack[:len(stack) - 1]
            }
        }
    }
    //check remaining in stack
    starCount := 0
    for i:=len(stack) - 1 ; i>=0 ; i-- {
        if stack[i] == "*" {
            starCount++
        } else if stack[i] == "(" {
            starCount--
        }
        if starCount < 0 {
            return false
        }
    }

    return true
}

func main() {
    var inputString string
    inputString = "(*)"
    fmt.Println(checkValidString(inputString))
}
