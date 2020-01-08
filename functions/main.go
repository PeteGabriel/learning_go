package main

func main() {

}


/**
Write a function that calculates the average of a float64 slice.
 */
func GetAvg(elems []float64) float64{
	var sum float64
	for _, e := range elems {
		sum = sum + e
	}
	return sum / float64(len(elems))
}

