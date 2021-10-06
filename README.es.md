# gcl

Esta biblioteca permite generar logs compatibles con la API de Google Cloud Logging que pueden ser usados en código implementado en Cloud Functions, Cloud Run, o Kubernetes.

## Métodos
```go
// Causa del error.
gcl.Cause(err)

// Campos personalizados con valores arbitrarios.
// Aparecerán en el registro de Cloud Logging en «jsonPayload».
gcl.Field("usuario", 573298)
gcl.Field("rol", "supervisor")

// Campos personalizados de etiquetas.
// Aparecerán en el registro de Cloud Logging en «labels».
gcl.Label("aplicacion", "mi-aplicacion")

// Descripción del error con nivel de severidad.
// Este debe ser un mensaje apto para responder al usuario final,
// como se verá más adelante, si se usa en peticiones HTTP.
gcl.Debug("ups")
gcl.Info("ups")
gcl.Notice("ups")
gcl.Warning("ups")
gcl.Error("ups")
gcl.Critical("ups")
gcl.Alert("ups")
gcl.Emergency("ups")

// Establecer código de estado de una petición HTTP, y datos de la petición.
gcl.HTTPStatus(http.StatusNotFound)
gcl.HTTPRequest(r)

// Todos los métodos anteriores pueden ser encadenados:
gcl.Cause(err).Field("usuario", 573298).Warning("usuario no encontrado").HTTPStatus(http.StatusNotFound)

// Los métodos anteriores inicializan la entrada del log, pero no generan el log, no escriben a stdout/stderr.
// Para escribir el log a stdout, se usa Log():
gcl.Cause(err).Field("usuario", 573298).Warning("usuario no encontrado").HTTPStatus(http.StatusNotFound).Log()

// Para retornar un error y darle posterior manejo al error, se usa Err() (ver ejemplo a continuación):
gcl.Cause(err).Field("usuario", 573298).Warning("usuario no encontrado").HTTPStatus(http.StatusNotFound).Err()

```

## Manejo de errores en peticiones HTTP:

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
	http.HandleFunc("/usuarios", usuarios)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	datos, err := obtenerUsuario(r.FormValue("usuario"))
	if err != nil {
		res, cod := manejarError(r, err)
		http.Error(w, res, cod)
		return
	}

	fmt.Fprint(w, datos)
}

func obtenerUsuario(u string) ([]byte, error) {
	if u == "" {
		return nil, gcl.Notice("no se recibió el usuario").HTTPStatus(http.StatusBadRequest).Err()
	}

	var usuario = struct {
		nombre string
	}{}

	err := db.QueryRow("SELECT nombre FROM usuarios WHERE id = $1", u).Scan(&usuario)
	if err == sql.ErrNoRows {
		return nil, gcl.Cause(err).Notice("usuario no encontrado").HTTPStatus(http.StatusNotFound).Err()
	} else if err != nil {
		return nil, gcl.Cause(err).Field("usuario", u).Error("ocurrió un error").HTTPStatus(http.StatusInternalServerError).Err()
	}

	b, err := json.Marshal(usuario)
	if err != nil {
		return nil, gcl.Cause(err).Field("usuario", u).Error("ocurrió un error").HTTPStatus(http.StatusInternalServerError).Err()
	}

	return b, nil
}

func manejarError(r *http.Request, e error) (string, int) {
	// Se espera un error *gcl.Err. Si no lo es, se usarán estos parámetros predeterminados:
	msj := "algo salió mal"
	cod := http.StatusInternalServerError
	res := fmt.Sprintf(`{"error":"%s"}`, msj)

	err, ok := e.(*gcl.Err)
	if ok {
		// Retorna el código establecido previamente, o 500 si no se estableció ninguno:
		cod = err.HTTPStatus()

		// Retorna el mensaje establecido en una de las ocho funciones con nivel de severidad (Debug, Info, Error, etc.), o msj si no se estableció ninguno.
		msj = err.Message(msj)

		res = fmt.Sprintf(`{"error":"%s"}`, msj)

		// Establecer datos de la petición HTTP:
		err.HTTPRequest(r)

		// Registrar log en Cloud Logging:
		err.Log()
	}

	return res, cod
}

```
