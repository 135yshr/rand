package randstr

// NewIDGenerator returns a new ID generator.
//
// The ID generator creates a string using alphanumeric characters.
func NewIDGenerator() Generator {
	return newStringGenerator(lowercase + uppercase + numbers)
}
