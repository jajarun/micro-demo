package main

import (
	"fmt"
	"go.uber.org/dig"
)

type ConfigTest struct {
	Name string
}

type GuestInterface interface {
	run()
}

type Guest struct {
	config ConfigTest
}

func (g *Guest) run() {
	fmt.Println("11111")
	fmt.Println(g.config.Name, " run")
}

func newGuest(c ConfigTest) Guest {
	return Guest{config: c}
}

func main() {
	container := dig.New()
	container.Provide(func() ConfigTest { return ConfigTest{Name: "test"} })
	container.Provide(func(c ConfigTest) GuestInterface {
		g := new(Guest)
		g.config = c
		return g
	})
	//container.Invoke(func(c ConfigTest) {
	//	fmt.Println(c.Name)
	//})
	err := container.Invoke(func(g GuestInterface) {
		g.run()
	})
	if err != nil {
		fmt.Println(err)
	}
}
