package main

import (
	"syscall"
	"testing"
	"time"
)

// basic functionality test
func TestNumberProducer(t *testing.T) {
	testChan := make(chan int)

	go numberProducer(time.Second, testChan, time.Millisecond)
	result := <-testChan

	if !(result >= 1 && result <= 100) {
		t.Fatalf(
			"Invalid integer received, numbers must be between 1 and 100 inclusive.  Got %d",
			result)
	}

}

// should run for one second and generate two results
func TestNumberProducerOneSecond(t *testing.T) {
	testChan := make(chan int)
	interval, _ := time.ParseDuration(".5s")

	go numberProducer(interval, testChan, time.Second)

	var result []int

	for {
		number, ok := <-testChan
		if !ok {
			break
		}
		result = append(result, number)
	}

	if len(result) != 2 {
		t.Fatalf("Wrong number of integers generated.  Expected two  Output: %v", result)
	}
}

// should run for ten seconds and generate twenty results
func TestNumberProducerTenSeconds(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	testChan := make(chan int)
	interval, _ := time.ParseDuration(".5s")

	go numberProducer(interval, testChan, 10*time.Second)

	var result []int

	for {
		number, ok := <-testChan
		if !ok {
			break
		}
		result = append(result, number)
	}

	if len(result) != 20 {
		t.Fatalf("Wrong number of integers generated.  Expected 20  Output: %v (%d elements)", result, len(result))
	}
}

// numberProducer should stop and clean up on a sigint signal
func TestNumberProducerSIGINT(t *testing.T) {
	testChan := make(chan int)

	go numberProducer(time.Second, testChan, 5*time.Second)

	_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)

	time.Sleep(time.Millisecond)

	// empty the int we generated when we first started up
	<-testChan
	// verify that the channel closed out
	_, ok := <-testChan
	if ok {
		close(testChan)
		t.Fatalf("Channel is still open, NumberProducer didn't halt as expected")
	}
}

// numberProducer should stop and clean up on a sigint signal
func TestNumberProducerSIGTERM(t *testing.T) {
	testChan := make(chan int)

	go numberProducer(time.Second, testChan, 5*time.Second)

	_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	time.Sleep(time.Millisecond)

	// empty the int we generated when we first started up
	<-testChan
	// verify that the channel closed out
	_, ok := <-testChan
	if ok {
		close(testChan)
		t.Fatalf("Channel is still open, NumberProducer didn't halt as expected")
	}
}
