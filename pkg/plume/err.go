package plume

type PlumeError struct {
	Reason string
}

func (err PlumeError) Error() string {
	return err.Reason
}
