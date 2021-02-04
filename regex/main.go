package main

import (
	"fmt"
)

func main() {

	fmt.Println(re.MatchString("q:dafas"), "q:")
	fmt.Println(re.MatchString("Q:fadasa"), "Q:")
	fmt.Println(re.MatchString("Ｑ:das"), "Ｑ:")
	fmt.Println(re.MatchString("Ｑ：asda"), "Ｑ：")
	fmt.Println(re.MatchString("aq：das"), "q：")
	fmt.Println(re.MatchString("Q：asd"), "Q：")
}
