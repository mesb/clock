package main

/*
  An implementation of a toy clock using recursion, and go concurrency patterns

*/


import (
	"fmt"
	"time"
)

/*
  tick: Actual ticking of a clock. It sends its ticks via a channel to the
  counter process, that does clock arithmetic.
*/

func tick(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "tock"
	tick(c)
}

/*
  updateCount: receives a tick, updates the seconds of the counter object, and
  send the counter via a channel tot he printer.
*/
func updateCount(tick chan string, ctr *Counter, ctrChan chan Counter) {
	t := <-tick
	if t != "" {
		ctr.plusSec()
		ctrChan <- *ctr
	}
	updateCount(tick, ctr, ctrChan)
}


/*
  Counter: represents the hour, minute, and seconds of a clock
*/
type Counter struct {
	hh, mm, ss int
}

/*
 plusSec: Increments seconds of a counter object
*/
func (c *Counter) plusSec() {
	if c.ss == 59 {
		c.plusMin()
		c.ss = 0
	}
	c.ss++
}

/*
  plusMin: increments the minutes of a counter object
*/
func (c *Counter) plusMin() {
	if c.mm == 59 {
		c.PlusHr()
		c.mm = 0
	}
	c.mm++
}

/*
  plusHr: Incremenets the hours of a counter object
*/
func (c *Counter) PlusHr() {
	if c.hh == 23 {
		c.hh = 0
	}
}

/*
  String representation of a counter, which is eventually a clock display
*/
func (c Counter) String() string{
	return fmt.Sprintf("%.2d : %.2d : %.2d", c.hh, c.mm, c.ss)
}

/*
   printer: receives a counter object via a channel and displays it, thereby
   showing the clock/counter
*/
func printer(C chan Counter) {
	c := <-C
	fmt.Printf("%s \r", c)
	printer(C)
}

/*
  main():
  1. Create a counter object
  2. create a channel of type string, for sending and receiving ticks
  3. Create a Counter channel for sending and receiving counter objects
  4. Concurrently run counter/clock processes
  5. Block the main function with an input operation.
*/
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
