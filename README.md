# clock
A simple clock application using go concurrency patterns

# Concurrency
"Concurrency is a composition of independently executing things"
                 -- Rob Pike, Concurrency is not parallism

# About this Application
This application is a composition of independently executing processes that
collectively simulate a toy clock. Its purpose is
investigate and enable me understand go's concurrency primitives and patterns.

# Technique
## Recursion
Recursion, in this application, is used to describe processes that run forever.

tick(c chan string)
updateCount(tick chan string, ctr *Counter, ctrChan chan Counter)

2 recursive functions, running forever. The tick function is the actual "tick"
of the clock. It constantly ticks and sends the tick to the updateCount function
that concurrently does the necessary clock arithmetic.

## Channels
Channels are used as pipelines for sending and receiving messages. Go provides
channels.

There is one channel that sends ticks from the tick function to the
updateCounter function.

The updateCounter also has a Counter channel that is used to send updated
counters to the printer function, which then prints the clock

## Functions
Functions describe the processes that collectively and concurrently implement
the toy clock.

There are 3 main processes in this application.

1. tick(c chan string):
tick describes the ticking of a clock

2. updateCounter(tick chan string, ctr *Counter, ctrChan chan Counter):
receives a tick, a pointer to the instantiated counter object, and a channel of
type Counter. It increments seconds when it receives ticks.

## Data Structures
Counter is the main data structure. It represents the hour, minutes, and
seconds of a clock.

Also, it has methods that increment seconds, minutes, and hours.


# Installation
1. go get -u github.com/mesb/clock
2. cd /path/to/clock --> Definitely in your workspace
3. go build
4. ./clock

# References
Communicating Sequential Processes, C.A.R Hoare

Concurrency is Not Parallelism, Rob Pike
