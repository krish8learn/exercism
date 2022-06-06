package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	// panic("Implement the ConcurrentFrequency function")
	ret := FreqMap{}
	//creating a channel which type is Freqmap//
	freqChannel := make(chan FreqMap)

	//create anonymous function to receive the result from the Frqunecy function//
	callFunc := func(s string) {
		freqChannel <- Frequency(s)
	}

	//lets call Frequency function through go routines for each text
	for _, text := range l {
		go callFunc(text)

		//insert the values, key from the channel
		for key, value := range <-freqChannel {
			ret[key] += value
		}
	}

	return ret
}
