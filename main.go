package main

import (
	"example/producer-consumer-project/consumer"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var interval time.Duration
	var concurrency int

	/* N.B. there is room for improvement here.  We should either prevent a multi-minute or longer run (with explanation)
	or teach the program how to extend the timeout in the case that is selected.
	The correct decision for how to proceed would depend on the intended use of the system.
	*/
	flag.DurationVar(&interval,
		"interval", time.Second,
		"Specify the interval for how often a new random number is generated.  Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\".")

	flag.IntVar(&concurrency,
		"concurrency", 1,
		"Specify the number of concurrent consumers of random numbers.")

	flag.Parse()

	// set up a channel for our producer and consumers
	// set the buffer size to the max amount possible given the rate and max duration, we shouldn't need it but don't want to block anything.
	numberChannel := make(chan int, (60 / interval))

	var wg sync.WaitGroup
	wg.Add(concurrency)

	go numberProducer(interval, numberChannel, time.Minute)
	for i := 0; i < concurrency; i++ {
		go consumer.ProcessNumbers(numberChannel, &wg)
	}

	wg.Wait()
}

/* N.B. if we want to be any more complex with the data generated we should probably break
this out into its own helper.  For this implementation it would probably be overkill*/

// numberProducer will generate random numbers at the interval provided
// and feed them to a channel for consumers
func numberProducer(interval time.Duration, channel chan int, timeout time.Duration) {
	// make sure the channel gets cleaned up when we are done sending
	defer close(channel)

	// make sure we're actually random
	rand.Seed(time.Now().UnixNano())

	// listen for SIGINT or SIGTERM so we can exit gracefully
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGTERM, syscall.SIGINT)

	var sleepTime time.Duration
	if interval < time.Millisecond {
		sleepTime = interval
	} else {
		sleepTime = time.Millisecond
	}

	stop_time := time.Now().Add(timeout)
	var lastSig time.Time

	for time.Now().Before(stop_time) {

		select {
		case <-osSig:
			return
		default:
			//make sure its time to generate a new number
			nextSignalTime := lastSig.Add(interval)
			if time.Now().After(nextSignalTime) {
				// get a new random number and put it into the channel
				integer := rand.Intn(100) + 1
				channel <- integer
				lastSig = time.Now()
			}

			time.Sleep(sleepTime)
		}
	}
}
