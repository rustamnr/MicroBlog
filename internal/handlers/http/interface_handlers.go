package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPHandler interface is used to create handlers for each model
type HTTPHandler interface {
	RegisterRouters(router *mux.Router)
	sendError(w http.ResponseWriter, message string, statusCode int)
}
