package main

import (
	"fmt"
	"time"
)


func tick(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "tock"
	tick(c)
}

func updateCount(tick chan string, ctr *Counter, ctrChan chan Counter) {
	t := <-tick
	if t != "" {
		ctr.plusSec()
		ctrChan <- *ctr
	}
	updateCount(tick, ctr, ctrChan)
}



type Counter struct {
	hh, mm, ss int
}


func (c *Counter) plusSec() {
	if c.ss == 59 {
		c.plusMin()
		c.ss = 0
	}
	c.ss++
}

func (c *Counter) plusMin() {
	if c.mm == 59 {
		c.PlusHr()
		c.mm = 0
	}
	c.mm++
}

func (c *Counter) PlusHr() {
	if c.hh == 23 {
		c.hh = 0
	}
}

func (c Counter) String() string{
	return fmt.Sprintf("%.2d : %.2d : %.2d", c.hh, c.mm, c.ss)
}

func printer(C chan Counter) {
	c := <-C
	fmt.Printf("%s \r", c)
	printer(C)
}

func main() {
	ctr := &Counter{0, 0, 0}
	var c chan string = make(chan string)
	var ctrChan chan Counter = make(chan Counter)

	go tick(c)
	go updateCount(c, ctr, ctrChan)
	go printer(ctrChan)

	var input string
	fmt.Scanln(&input)
}
