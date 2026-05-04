package internal

func CalculatePi(start int, routineNum int, stopSignal chan struct{}, resultChannel chan float64) {
	var sum float64
	for {
		select {
		case <-stopSignal:
			resultChannel <- sum
			return
		default:
			sign := 0.0
			if start%2 == 0 {
				sign = 1
			} else {
				sign = -1
			}
			element := sign / float64(2*start+1)
			sum += element
			start += routineNum
		}
	}
}
