package output

// from http://stackoverflow.com/a/16931348/2688600
func FanOut(ch <-chan int32, size, lag int32) []chan int32 {
	cs := make([]chan int32, size)
	for i, _ := range cs {
		// The size of the channels buffer controls how far behind the recievers
		// of the fanOut channels can lag the other channels.
		cs[i] = make(chan int32, lag)
	}
	go func() {
		for i := range ch {
			for _, c := range cs {
				c <- i
			}
		}
		for _, c := range cs {
			// close all our fanOut channels when the input channel is exhausted.
			close(c)
		}
	}()
	return cs
}
