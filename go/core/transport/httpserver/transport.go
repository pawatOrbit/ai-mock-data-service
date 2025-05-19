package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"reflect"
	"time"

	"github.com/pawatOrbit/ai-mock-data-service/go/core/transport"
)

func NewTransport[T, R any](req T, endpoint func() Endpoint[T, R], middlewares ...transport.EndpointMiddleware[T, R]) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		newReq := deepCopy(req)

		var (
			ctx            = r.Context()
			httpStatusCode = http.StatusOK
			method         = r.Method
			path           = r.URL.Path
			resp           R
			serviceError   error
			elapsedTime    time.Duration
		)

		requestBody, err := readRequestBody(r)
		if err != nil {
			fmt.Println("Error reading request body")
			HandleInternalServerError(w, http.StatusBadRequest)
			return
		}

		if len(requestBody) == 0 {
			requestBody = []byte("{}")
		}

		err = json.Unmarshal(requestBody, &newReq)
		if err != nil {
			fmt.Println("Error unmarshalling request body")
			HandleInternalServerError(w, http.StatusBadRequest)
			return
		}

		startTime := time.Now()
		resp, serviceError = endpoint()()(r.Context(), newReq)
		elapsedTime = time.Since(startTime)

		if serviceError != nil {
			httpStatusCode = http.StatusInternalServerError
			HandleInternalServerError(w, httpStatusCode)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpStatusCode)
			json.NewEncoder(w).Encode(resp)

			slog.InfoContext(ctx, "HTTP Request",
				slog.String("method", method),
				slog.String("path", path),
				slog.Duration("elapsed_time", elapsedTime),
			)
			return
		}

	}
}

func deepCopy[T any](src T) T {
	return reflect.New(reflect.TypeOf(src).Elem()).Interface().(T)
}

func readRequestBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}
