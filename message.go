package gcl

// Debug sets an user message with DEBUG severity.
func Debug(msg string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.setSeverityAndMessage("DEBUG", msg)
	return e
}

// Info sets an user message with INFO severity.
func Info(msg string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.setSeverityAndMessage("INFO", msg)
	return e
}

// Notice sets an user message with NOTICE severity.
func Notice(msg string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.setSeverityAndMessage("NOTICE", msg)
	return e
}

// Warning sets an user message with WARNING severity.
func Warning(msg string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.setSeverityAndMessage("WARNING", msg)
	return e
}

// Error sets an user message with ERROR severity.
func Error(msg string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.setSeverityAndMessage("ERROR", msg)
	return e
}

// Critical sets an user message with CRITICAL severity.
func Critical(msg string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.setSeverityAndMessage("CRITICAL", msg)
	return e
}

// Alert sets an user message with ALERT severity.
func Alert(msg string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.setSeverityAndMessage("ALERT", msg)
	return e
}

// Emergency sets an user message with EMERGENCY severity.
func Emergency(msg string) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	e.setSeverityAndMessage("EMERGENCY", msg)
	return e
}

// Debug sets an user message with DEBUG severity.
func (e *Entry) Debug(msg string) *Entry {
	e.setSeverityAndMessage("DEBUG", msg)
	return e
}

// Info sets an user message with INFO severity.
func (e *Entry) Info(msg string) *Entry {
	e.setSeverityAndMessage("INFO", msg)
	return e
}

// Notice sets an user message with NOTICE severity.
func (e *Entry) Notice(msg string) *Entry {
	e.setSeverityAndMessage("NOTICE", msg)
	return e
}

// Warning sets an user message with WARNING severity.
func (e *Entry) Warning(msg string) *Entry {
	e.setSeverityAndMessage("WARNING", msg)
	return e
}

// Error sets an user message with ERROR severity.
func (e *Entry) Error(msg string) *Entry {
	e.setSeverityAndMessage("ERROR", msg)
	return e
}

// Critical sets an user message with CRITICAL severity.
func (e *Entry) Critical(msg string) *Entry {
	e.setSeverityAndMessage("CRITICAL", msg)
	return e
}

// Alert sets an user message with ALERT severity.
func (e *Entry) Alert(msg string) *Entry {
	e.setSeverityAndMessage("ALERT", msg)
	return e
}

// Emergency sets an user message with EMERGENCY severity.
func (e *Entry) Emergency(msg string) *Entry {
	e.setSeverityAndMessage("EMERGENCY", msg)
	return e
}

// setSeverityAndMessage sets the severity and message.
func (e *Entry) setSeverityAndMessage(severity string, message string) {
	e.Severity = severity
	e.Message = message
}
