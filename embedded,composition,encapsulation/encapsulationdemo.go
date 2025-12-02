package main

import "fmt"

type Counter struct {
	count int
}

func NewCounter() *Counter { return &Counter{} }

func (c *Counter) inc()       { c.count++ }
func (c *Counter) Dec()       { c.count-- }
func (c *Counter) value() int { return c.count }

func main() {
	c := NewCounter()
	c.inc()
	c.inc()
	c.inc()
	c.Dec()
	c.inc()
	c.inc()
	c.Dec()
	c.inc()
	c.Dec()
	fmt.Println(c.value())

}
