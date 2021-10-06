package gcl

import (
	"fmt"
)

// Label sets a key-value label.
func Label(key string, value interface{}) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	return e.Label(key, value)
}

// Label sets a key-value label.
func (e *Entry) Label(key string, value interface{}) *Entry {
	if e.Labels == nil {
		e.Labels = map[string]string{}
	}

	e.Labels[key] = fmt.Sprintf("%v", value)

	return e
}
