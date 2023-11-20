package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.URL.Query().Get("nome")

	res, err := rt.db.GetName()
	w.Header().Set("content-type", "text/plain")
	json.NewEncoder(w).Encode(err)
	json.NewEncoder(w).Encode(res)
}
