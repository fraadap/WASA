package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u user
	err = json.Unmarshal(body, &u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := rt.db.Login(u.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(id)

}
