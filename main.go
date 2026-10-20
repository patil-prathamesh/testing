package main

import (
	"fmt"
	"os"
	"time"

	"p.com/hello/mocking"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func walk(x interface{}, fn func(string)) {
	fmt.Println("x")
	fn(x.(string))
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)

	fn := func(s string) {
		fmt.Println(s)
	}

	walk(struct{ s string }{"patil"}, fn)
}
