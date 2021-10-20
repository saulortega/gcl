package gcl

// Trace creates a new Entry and sets a log trace.
func Trace(trace string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.TraceLog = trace
	return e
}

// Trace sets a log trace for the Entry.
func (e *Entry) Trace(trace string) *Entry {
	e.TraceLog = trace
	return e
}
