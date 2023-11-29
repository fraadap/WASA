package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/fraadap/WASA/service/api/reqcontext"
	"github.com/fraadap/WASA/service/structs"
	"github.com/julienschmidt/httprouter"
)

// Posts a new ban
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, er := io.ReadAll(r.Body)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var b structs.Ban
	err0 := json.Unmarshal(body, &b)
	if b.TimeStamp == "" {
		b.TimeStamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	} else if len(b.TimeStamp) != 20 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err0 != nil || b.Banned == 0 || b.Banned == id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b.UserID = id

	token := getToken(r.Header.Get("Authorization"))
	if b.UserID != token {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var err1 error
	b.BanID, err1 = rt.db.NewBan(id, b.Banned, b.TimeStamp)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

// Removes a ban
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error
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

	// ban id in input da cambiare
	banID, err := strconv.Atoi(ps.ByName("banID"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err1 := rt.db.DeleteBan(id, banID)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
