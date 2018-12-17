package proto

type (
	// Message is contains a payload and the source that generated the payload
	Message struct {
		Source  string
		Payload []int
	}
)
