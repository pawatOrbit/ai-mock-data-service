package middleware_httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/pawatOrbit/ai-mock-data-service/go/core/transport/httpserver"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	resp := httpserver.ModelResp{
		Status:  http.StatusNotFound,
		Message: "Not Found",
	}

	json.NewEncoder(w).Encode(resp)
}
