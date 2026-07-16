package internal_test

import (
	"math"
	intnl "solvery/04_lesson/internal"
	"testing"
	"time"
)

func TestCalculatePi(t *testing.T) {
	stopChannel := make(chan struct{})
	numberOfRoutines := 3
	var sum float64
	resultChannel := make(chan float64, numberOfRoutines)

	for i := 0; i < numberOfRoutines; i++ {
		go intnl.CalculatePi(i, numberOfRoutines, stopChannel, resultChannel)
	}

	time.Sleep(5 * time.Second)
	close(stopChannel)

	for i := 0; i < numberOfRoutines; i++ {
		sum += <-resultChannel
	}
	sum = sum * 4

	if math.Abs(sum-math.Pi) > 1e-6 {
		t.Errorf("CalculatePi returned wrong result")
	}
}
