package gcl

import (
	"runtime"
	"strconv"
)

type SourceLocation struct {
	File     string `json:"file,omitempty"`
	Line     string `json:"line,omitempty"`
	Function string `json:"function,omitempty"`
}

// setSourceLocation sets the source location where the first method is called from.
func (e *Entry) setSourceLocation() {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		var function string
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			function = fn.Name()
		}

		e.SourceLocation = &SourceLocation{
			File:     file,
			Line:     strconv.Itoa(line),
			Function: function,
		}
	}
}
