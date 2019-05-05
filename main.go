package main

import (
	"fmt"
	"lab/service"
)

func main() {
	d := service.NewDB()

	g := service.NewGreeter(d, "en")
	fmt.Println(g.Greet())
	fmt.Println(g.GreetInDefaultMessage())

	g = service.NewGreeter(d, "es")
	fmt.Println(g.Greet())

	g = service.NewGreeter(d, "random")
	fmt.Println(g.Greet())
}
