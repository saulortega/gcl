package gcl

// Cause sets an error cause.
// If err is *Err type, the initial Entry from that Err is returned.
func Cause(err error) *Entry {
	if E, ok := err.(*Err); ok {
		return E.entry
	}

	e := &Entry{}
	e.setSourceLocation()
	e.CauseErr = err
	return e
}

// Cause sets an error cause.
// If err is *Err type, the initial Entry from that Err is returned.
func (e *Entry) Cause(err error) *Entry {
	if E, ok := err.(*Err); ok {
		return E.entry
	}

	e.CauseErr = err
	return e
}
