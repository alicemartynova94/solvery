package main

import (
	"fmt"
	"os"
	"os/signal"
	"solvery/lesson_four/internal"
)

func main() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	stopChannel := make(chan struct{})
	numberOfRoutines := 3
	var sum float64

	resultChannel := make(chan float64, numberOfRoutines)

	for i := 0; i < numberOfRoutines; i++ {
		go internal.CalculatePi(i, numberOfRoutines, stopChannel, resultChannel)
	}

	<-signalChannel
	close(stopChannel)

	for i := 0; i < numberOfRoutines; i++ {
		sum += <-resultChannel
	}

	fmt.Printf("\nResult: %.10f\n", sum*4)
}
