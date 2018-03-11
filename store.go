package flux

type StoreInterface interface {
	// Handle consumes the action in the payload. If the operation was performed syncronously,
	// return true. If not, return false and close the payload.Done channel when finished.
	Handle(payload *Payload) (finished bool)
}
