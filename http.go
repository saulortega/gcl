package gcl

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Request represents an HTTP Request.
type Request struct {
	RequestMethod string `json:"requestMethod,omitempty"`
	RequestURL    string `json:"requestUrl,omitempty"`
	UserAgent     string `json:"userAgent,omitempty"`
	RemoteIP      string `json:"remoteIp,omitempty"`
	Referer       string `json:"referer,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	ResponseSize  string `json:"responseSize,omitempty"`
	Latency       string `json:"latency,omitempty"`
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

// HTTPResponseSize sets the size of the HTTP response.
func HTTPResponseSize(size int) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	return e.HTTPResponseSize(size)
}

// HTTPLatency sets the latency of the HTTP request processing.
func HTTPLatency(duration time.Duration) *Entry {
	e := &Entry{}
	e.setSourceLocation()
	return e.HTTPLatency(duration)
}

// HTTPRequest sets a HTTP Request.
func (e *Entry) HTTPRequest(r *http.Request) *Entry {
	if e.Request == nil {
		e.Request = &Request{}
	}

	e.Request.RequestMethod = r.Method
	e.Request.RequestURL = r.URL.String()
	e.Request.UserAgent = r.UserAgent()
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

// HTTPResponseSize sets the size of the HTTP response.
func (e *Entry) HTTPResponseSize(size int) *Entry {
	if e.Request == nil {
		e.Request = &Request{}
	}

	e.Request.ResponseSize = strconv.Itoa(size)

	return e
}

// HTTPLatency sets the latency of the HTTP request processing.
func (e *Entry) HTTPLatency(duration time.Duration) *Entry {
	if e.Request == nil {
		e.Request = &Request{}
	}

	e.Request.Latency = fmt.Sprintf("%.9fs", duration.Seconds())

	return e
}
