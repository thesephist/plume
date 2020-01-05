package plume

// Error represents any error originating
// from unexpected states in the Plume server.
type Error struct {
	reason string
}

func (err Error) Error() string {
	return err.reason
}
