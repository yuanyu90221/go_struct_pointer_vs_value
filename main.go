package main

import (
	"fmt"
	"time"
)
type car struct {
	name string
}

type email struct {
	from string
	to string
}

func(c car) SetName01(s string) {
	fmt.Printf("1. address: %p\n", &c);
	c.name = s
}

func(c *car) SetName02(s string){
	c.name = s
}
func(e email) From(s string) email{
	e.from = s
	return e
}
func(e email) To(s string) email{
	e.to = s
	return e
}

func(e email) Send() {
	fmt.Printf("From: %s, To: %s\n", e.from, e.to)
}
func main() {
	c := &car{}
	fmt.Printf("0 address: %p\n", c);
	c.SetName01("foo")
	fmt.Println(c.name)
	c.SetName02("bar")
	fmt.Println(c.name)

	
	e := &email{}
	for i:=1; i<=10; i++ {
		go func(i int) {
			// e := &email{}
			e.From(fmt.Sprintf("example%02d@example.com", i)).To(fmt.Sprintf("example%02d@example.com", i+1)).Send()
			// e.To(fmt.Sprintf("example%02d@example.com", i+1))
			// e.Send()
		}(i)
	}
	time.Sleep(1 * time.Second)
}