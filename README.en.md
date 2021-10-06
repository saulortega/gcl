# gcl

This library allows to generate logs compatible with the Google Cloud Logging API that can be used in code implemented in Cloud Functions, Cloud Run, or Kubernetes.

## Methods
```go
// error cause.
gcl.Cause(err)

// Custom fields with arbitrary values.
// Will appear in the entry of Cloud Logging in «jsonPayload».
gcl.Field("user", 573298)
gcl.Field("role", "supervisor")

// Custom fields for labels.
// Will appear in the entry of Cloud Logging in «labels».
gcl.Label("app", "my-app")

// Description of the error with severity.
// This must be a message suitable for be replied to the final user
// when it is used in HTTP requests.
gcl.Debug("oops")
gcl.Info("oops")
gcl.Notice("oops")
gcl.Warning("oops")
gcl.Error("oops")
gcl.Critical("oops")
gcl.Alert("oops")
gcl.Emergency("oops")

// Set HTTP status code and data from HTTP request.
gcl.HTTPStatus(http.StatusNotFound)
gcl.HTTPRequest(r)

// All the above methods can be chained:
gcl.Cause(err).Field("user", 573298).Warning("user not found").HTTPStatus(http.StatusNotFound)

// The above methods initialize the entry of log, but does not write the log.
// To write the log to stdout, use Log():
gcl.Cause(err).Field("user", 573298).Warning("user not found").HTTPStatus(http.StatusNotFound).Log()

// To return an error and handle it later, use Err():
gcl.Cause(err).Field("user", 573298).Warning("user not found").HTTPStatus(http.StatusNotFound).Err()

```

## Error handling in HTTP requests:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/saulortega/gcl"
)

func main() {
	http.HandleFunc("/users", users)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func users(w http.ResponseWriter, r *http.Request) {
	data, err := getUser(r.FormValue("user"))
	if err != nil {
		res, cod := handleError(r, err)
		http.Error(w, res, cod)
		return
	}

	fmt.Fprint(w, data)
}

func getUser(u string) ([]byte, error) {
	if u == "" {
		return nil, gcl.Notice("empty user").HTTPStatus(http.StatusBadRequest).Err()
	}

	var user = struct {
		name string
	}{}

	err := db.QueryRow("SELECT name FROM users WHERE id = $1", u).Scan(&user)
	if err == sql.ErrNoRows {
		return nil, gcl.Cause(err).Notice("user not found").HTTPStatus(http.StatusNotFound).Err()
	} else if err != nil {
		return nil, gcl.Cause(err).Field("user", u).Error("something went wrong").HTTPStatus(http.StatusInternalServerError).Err()
	}

	b, err := json.Marshal(user)
	if err != nil {
		return nil, gcl.Cause(err).Field("user", u).Error("something went wrong").HTTPStatus(http.StatusInternalServerError).Err()
	}

	return b, nil
}

func handleError(r *http.Request, e error) (string, int) {
	// An *gcl.Err is expected. If not, this parameters will be used:
	msg := "something went wrong"
	cod := http.StatusInternalServerError
	res := fmt.Sprintf(`{"error":"%s"}`, msg)

	err, ok := e.(*gcl.Err)
	if ok {
		// Returns the code previously setted, or 500 if not was setted.
		cod = err.HTTPStatus()

		// Returns the setted message in one of the eight functions with severity (Debug, Info, Error, etc.), or msg if not was setted.
		msg = err.Message(msg)

		res = fmt.Sprintf(`{"error":"%s"}`, msg)

		// Set HTTP Request data:
		err.HTTPRequest(r)

		// Write log to stdout:
		err.Log()
	}

	return res, cod
}

```
