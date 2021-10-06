package gcl

// Cause sets an error cause.
func Cause(err error) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.CauseErr = err
	return e
}

// Cause sets an error cause.
func (e *Entry) Cause(err error) *Entry {
	e.CauseErr = err
	return e
}
