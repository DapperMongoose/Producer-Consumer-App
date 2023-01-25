package consumer

import (
	"sync"
	"testing"
)

// basic functionality
func TestReport(t *testing.T) {
	want := "{\"count\":100,\"average\":1}"
	output := report(100, 100)
	if output != want {
		t.Fatalf("expected %s but got %s", want, output)
	}
}

// if we had a zero sum somehow (likely by no data being sent to that worker)
// we should handle that gracefully
func TestReportZeroSum(t *testing.T) {
	want := "{\"count\":10,\"average\":0}"
	output := report(10, 0)
	if output != want {
		t.Fatalf("expected %s but got %s", want, output)
	}
}

// be extra sure that we handle zero values in both fields gracefully
func TestReportAllZeros(t *testing.T) {
	want := "{\"count\":0,\"average\":0}"
	output := report(0, 0)
	if output != want {
		t.Fatalf("expected %s but got %s", want, output)
	}
}

/*
N.B. It's hard to test stdout in go without bringing in more dependencies.
I would avoid doing it unless absolutely necessary and instead test the output that was being
set somewhere other than the stdout if it existed.

In that case we'd use a helper like this before each test to initialize the things we need
to test the ProcessNumbers function.
*/
func setupProcessNumbersTests(buffsize int) (chan int, *sync.WaitGroup) {
	testChan := make(chan int, buffsize)
	var wg sync.WaitGroup

	return testChan, &wg
}
