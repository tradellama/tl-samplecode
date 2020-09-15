package movingaverage

import "testing"

func TestSimpleMean(t *testing.T) {

	payload := []float64{
		2.0,
		3.0,
		4.0,
	}

	result, e := SimpleMean(payload)
	if e != nil {
		t.Errorf("failed to calculate average: %s", e)
	}
	if result != 3.0 {
		t.Errorf("\nexpected:\t3.0\ngot:\t\t\t%f", result)
	}
}

func TestSimpleMeanNil(t *testing.T) {

	payload := []float64{}

	_, e := SimpleMean(payload)
	if e == nil {
		t.Error("\nexpected:\tan error\ngot:\t\t\tnil")
	}
}

func TestStreamingMean(t *testing.T) {
	// Initialize our channels, to simulate a stream of data coming in
	// and a stream to get the results back in real time
	seed := make(chan float64)
	result := make(chan float64)

	go StreamingMean(seed, result)

	// send some data to the stream and read the result
	seed <- 1.0
	if x := <-result; x != 1.0 {
		t.Errorf("\nexpected:\t1.0\ngot:\t\t\t%f", x)
	}

	seed <- 2.0
	if x := <-result; x != 1.5 {
		t.Errorf("\nexpected:\t1.5\ngot:\t\t\t%f", x)
	}

	seed <- 3.0
	if x := <-result; x != 2.0 {
		t.Errorf("\nexpected:\t2.0\ngot:\t\t\t%f", x)
	}

	//close the stream to exit the goroutine
	close(seed)

}
