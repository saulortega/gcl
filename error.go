package gcl

import (
	"errors"
	"fmt"
	"strings"
)

type Err struct {
	entry *Entry
}

// Err returns a standard error interface from the Entry.
func (e *Entry) Err() Err {
	return Err{e}
}

// Error returns the json-structured error compatible with Google Logging.
func (e Err) Error() string {
	return e.entry.String()
}

// Log logs the error to stdout.
func (e *Err) Log() {
	fmt.Println(e.entry)
}

func (e Err) Is(target error) bool {
	return errors.Is(e.entry.causeErr, target)
}

func (e Err) As(target interface{}) bool {
	return errors.As(e.entry.causeErr, target)
}

func (e Err) Unwrap() error {
	return e.entry.causeErr
}

func (e Err) Entry() *Entry {
	return e.entry
}

// HTTPStatus returns the setted HTTP Status Code or 500 if not setted.
func (e *Err) HTTPStatus() int {
	if e.entry.Request == nil {
		return 500
	}

	if e.entry.Request.Status <= 0 {
		return 500
	}

	return e.entry.Request.Status
}

// Message returns the setted message or altMsg if not message is setted.
func (e *Err) Message(altMsg ...string) string {
	if e.entry.Message == "" && len(altMsg) > 0 {
		return strings.Join(altMsg, " - ")
	}

	return e.entry.Message
}
