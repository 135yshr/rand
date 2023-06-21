package randstr

// NewIDGenerator returns a new ID generator.
func NewIDGenerator() Generator {
	return newStringGenerator(lowercase + uppercase + numbers)
}
