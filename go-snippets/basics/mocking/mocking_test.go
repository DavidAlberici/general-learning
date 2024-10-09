package main

import (
	asserts "hello-world"
	"testing"
	"time"
)

type SpyCountdownOperations struct {
	Calls   []string
	Written []byte
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	s.Written = append(s.Written, p...)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("prints", func(t *testing.T) {
		spy := &SpyCountdownOperations{}

		Countdown(spy, spy)

		actualPrint := string(spy.Written)
		expectedPrint := `3
2
1
Go!`

		asserts.AssertStringEquals(actualPrint, expectedPrint, t)
	})

	t.Run("sleep between writes", func(t *testing.T) {
		spy := &SpyCountdownOperations{}

		Countdown(spy, spy)

		actualCalls := spy.Calls
		expectedCalls := []string{
			write, sleep,
			write, sleep,
			write, sleep,
			write,
		}

		asserts.AssertStringArraysEqual(actualCalls, expectedCalls, t)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
