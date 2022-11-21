package main

import (
	"os"

	"github.com/bietdoikiem/go-with-tdd/mocking"
)

func main() {
	// Default sleeper init
	defaultSleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, defaultSleeper)
}
