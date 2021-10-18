package gcl

import "errors"

// Cause sets an error cause.
// If err is *Err type, the initial Entry from that Err is returned.
func Cause(err error) *Entry {
	E := &Err{}
	if errors.As(err, E) {
		return E.entry
	} else if errors.As(err, &E) {
		return E.entry
	}

	e := &Entry{}
	e.setSourceLocation()
	e.causeErr = err
	if e.causeErr != nil {
		e.CauseErr = e.causeErr.Error()
	}

	return e
}

// Cause sets an error cause.
// If err is *Err type, the initial Entry from that Err is returned.
func (e *Entry) Cause(err error) *Entry {
	E := &Err{}
	if errors.As(err, E) {
		return E.entry
	} else if errors.As(err, &E) {
		return E.entry
	}

	e.causeErr = err
	e.CauseErr = ""
	if e.causeErr != nil {
		e.CauseErr = e.causeErr.Error()
	}

	return e
}
