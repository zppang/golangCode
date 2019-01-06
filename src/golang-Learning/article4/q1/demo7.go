package main

import (
	"flag"
	"fmt"
)

func main() {
	//方式一
	// name string
	//flag.StringVar(&name, "name", "everyone", "the Greeting Object")


	//方式二
	//var name = *flag.String("name", "everyone", "the Greeting Object")


	//方式三
	name := *flag.String("name", "everyone", "the Greeting Object")

	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
}
