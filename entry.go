package gcl

import (
	"encoding/json"
	"fmt"
)

type Entry struct {
	Message        string            `json:"message"`
	CauseErr       string            `json:"error,omitempty"`
	Severity       string            `json:"severity,omitempty"`
	Request        *Request          `json:"httpRequest,omitempty"`
	Labels         map[string]string `json:"logging.googleapis.com/labels,omitempty"`
	SourceLocation *SourceLocation   `json:"logging.googleapis.com/sourceLocation,omitempty"`
	causeErr       error
	fields         map[string]interface{}
}

// Log logs the entry to stdout.
func (e *Entry) Log() {
	fmt.Println(e)
}

// String returns the JSON representation of the entry in the format expected by Cloud Logging.
func (e Entry) String() string {
	out, err := json.Marshal(e)
	if err != nil {
		fmt.Printf("json.Marshal: %v\n", err)
		if e.causeErr != nil {
			fmt.Println(e.causeErr)
		}
		if e.Message != "" {
			fmt.Println(e.Message)
		}
	}

	out = e.appendFields(out)

	return string(out)
}

func (e *Entry) appendFields(out []byte) []byte {
	if len(e.fields) > 0 && len(out) > 0 {
		fields := map[string]string{}
		for key, value := range e.fields {
			fields[key] = fmt.Sprintf("%v", value)
		}

		outFields, err := json.Marshal(fields)
		if err != nil {
			fmt.Printf("json.Marshal: %v\n", err)
		} else {
			out = out[:len(out)-1]
			out = append(out, ',')
			out = append(out, outFields[1:len(outFields)-1]...)
			out = append(out, '}')
		}
	}

	return out
}
