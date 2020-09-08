package middlewares

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func RouterLogger(r *mux.Router) http.Handler{
	return handlers.LoggingHandler(os.Stdout, r)
}
