package movingaverage

import (
	"errors"
)

// SimpleMean calculates the mean for a given list of float64 values
func SimpleMean(x []float64) (float64, error) {
	if len(x) == 0 {
		return 0.0, errors.New("cannot calculate the average of an empty slice")
	}
	s := 0.0
	for _, v := range x {
		s += v
	}
	return s / float64(len(x)), nil
}

// StreamingMean calculates the mean for a given stream of float64 sent to a channel
func StreamingMean(in chan float64, out chan float64) {
	cnt := 0
	mean := 0.0
	for x := range in {
		cnt++
		differential := (x - mean) / float64(cnt)
		mean = mean + differential
		out <- mean
	}
}
