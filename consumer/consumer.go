package consumer

import (
	"encoding/json"
	"fmt"
	"sync"
)

type reportOutput struct {
	Count   int     `json:"count"`
	Average float64 `json:"average"`
}

func ProcessNumbers(channel chan int, wg *sync.WaitGroup) {
	// when this function ends for whatever reason, tell the wait group about it
	defer wg.Done()

	counter := 0
	sum := 0

	for {
		number, ok := <-channel
		if !ok {
			report(counter, sum)
			return
		}
		counter += 1
		sum += number
	}
}

func report(count int, sum int) string {
	var average float64
	if sum > 0 {
		average = float64(sum) / float64(count)
	} else {
		average = 0
	}

	output := reportOutput{
		Count:   count,
		Average: average,
	}

	formatted := encodeReport(output)

	fmt.Println(formatted)
	// return the value as well in case we care for testing or future expansion
	return formatted
}

func encodeReport(report reportOutput) string {
	bytes, err := json.Marshal(report)
	if err != nil {
		fmt.Printf("Error marshalling JSON to print report: %s\n", err)
		return ""
	}
	return (string(bytes))

}
