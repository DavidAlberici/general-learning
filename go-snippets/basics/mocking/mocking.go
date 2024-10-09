package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := range countdownStart {
		fmt.Fprintln(writer, countdownStart-i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	// argsWithProgramPath := os.Args
	args := os.Args[1:]

	sleepingDuration := 1 * time.Second

	if len(args) > 0 {
		sleepingSecondsArg, err := strconv.Atoi(args[0])
		if err == nil {
			sleepingDuration = time.Duration(sleepingSecondsArg) * time.Second
		}
	}

	sleeper := &ConfigurableSleeper{sleepingDuration, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
