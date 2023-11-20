package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// doLogin is
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte("Hello World!"))
}
