package main

// Logger provided
type Logger struct {
	timestampMap map[string]int
}

// Constructor provided
func Constructor() Logger {
	return Logger{timestampMap: make(map[string]int)}
}

// ShouldPrintMessage returns true if the message should be printed in the given timestamp,
// otherwise returns false.
// If this method returns false, the message will not be printed.
// The timestamp is in seconds granularity. */
func (logger *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	result := false
	if savedTime, isHere := logger.timestampMap[message]; isHere {
		if savedTime+10 <= timestamp {
			result = true
			// update map
			logger.timestampMap[message] = timestamp
		}
	} else {
		// not exist
		result = true
		// update map
		logger.timestampMap[message] = timestamp
	}

	// TODO clear old records by filtering timestamps

	return result
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

func main() {

	logger := Constructor()

	// logging string "foo" at timestamp 1, true
	println(logger.ShouldPrintMessage(1, "foo"))

	// logging string "bar" at timestamp 2, true
	println(logger.ShouldPrintMessage(2, "bar"))

	// logging string "foo" at timestamp 3, false
	println(logger.ShouldPrintMessage(3, "foo"))

	// logging string "bar" at timestamp 8, false
	println(logger.ShouldPrintMessage(8, "bar"))

	// logging string "foo" at timestamp 10, false
	println(logger.ShouldPrintMessage(10, "foo"))

	// logging string "foo" at timestamp 11, true
	println(logger.ShouldPrintMessage(11, "foo"))
}
