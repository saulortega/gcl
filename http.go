package gcl

import (
	"net/http"
)

// Request represents an HTTP Request.
type Request struct {
	RequestMethod string `json:"requestMethod,omitempty"`
	RequestUrl    string `json:"requestUrl,omitempty"`
	RemoteIP      string `json:"remoteIp,omitempty"`
	Referer       string `json:"referer,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	Status        int    `json:"status,omitempty"`
}

// HTTPRequest sets a HTTP Request.
func HTTPRequest(r *http.Request) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	return e.HTTPRequest(r)
}

// HTTPStatus sets the HTTP Status Code of an HTTP Request.
func HTTPStatus(status int) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	return e.HTTPStatus(status)
}

// HTTPRequest sets a HTTP Request.
func (e *Entry) HTTPRequest(r *http.Request) *Entry {
	if e.Request == nil {
		e.Request = &Request{}
	}

	e.Request.RequestMethod = r.Method
	e.Request.RequestUrl = r.URL.String()
	e.Request.RemoteIP = r.RemoteAddr
	e.Request.Referer = r.Header.Get("Referer")
	e.Request.Protocol = r.Proto

	return e
}

// HTTPStatus sets the HTTP Status Code of an HTTP Request.
func (e *Entry) HTTPStatus(status int) *Entry {
	if e.Request == nil {
		e.Request = &Request{}
	}

	e.Request.Status = status

	return e
}
