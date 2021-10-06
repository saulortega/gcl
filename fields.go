package gcl

// Field sets a key-value pair with relevant information.
func Field(key string, value interface{}) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	return e.Field(key, value)
}

// Field sets a key-value pair with relevant information.
func (e *Entry) Field(key string, value interface{}) *Entry {
	if e.fields == nil {
		e.fields = map[string]interface{}{}
	}

	e.fields[key] = value

	return e
}
