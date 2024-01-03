package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/fraadap/WASA/service/api/reqcontext"
	"github.com/fraadap/WASA/service/structs"
	"github.com/julienschmidt/httprouter"
)

// getUserProfile restituisce le foto dell'utente in ordine cronologico, quante foto ha, followers e following
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pr, err := rt.db.GetProfile(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	token := getToken(r.Header.Get("Authorization"))
	if token != id {
		pr.Bans = nil
	}
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w).Encode(pr)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := getToken(r.Header.Get("Authorization"))

	if id != token {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	body, er := io.ReadAll(r.Body)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u structs.User
	err0 := json.Unmarshal(body, &u)

	if err0 != nil || u.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err1 := rt.db.SetUsername(id, u.Username)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	u.Id = id

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	er1 := json.NewEncoder(w).Encode(u)
	if er1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (rt *_router) getUsernameByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username, err1 := rt.db.GetUsername(id)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	er1 := json.NewEncoder(w).Encode(username)
	if er1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
