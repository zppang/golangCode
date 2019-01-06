package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
}

/*func getTheFlag() *string {
	return  flag.String("name", "everyone", "The greeting object")
}*/


//上面函数的实现也可以是这样的。 在不改变主程序的情况下，便于对程序（如函数）的重构
func getTheFlag() *int {
	return  flag.Int("name", 1, "The greeting object")
}