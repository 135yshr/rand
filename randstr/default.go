package randstr

// NewDefaultGenerator returns a new Default generator.
func NewDefaultGenerator() Generator {
	return (newStringGenerator(lowercase + uppercase + numbers + symbols))
}
