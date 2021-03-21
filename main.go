package main

import (
	"os"
	"time"

	"github.com/FollowTheProcess/testygo/mocking"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &mocking.ConfigurableSleeper{
		Duration:  1 * time.Second,
		SleepFunc: time.Sleep,
	}
	mocking.Countdown(os.Stdout, sleeper)
}
