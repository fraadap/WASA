package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Posts a new ban
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var b ban
	err0 := json.Unmarshal(body, &b)
	if b.TimeStamp == "" {
		b.TimeStamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	} else if len(b.TimeStamp) != 20 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err0 != nil || b.UserIDBanned == 0 || b.UserIDBanned == id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var err1 error
	b.BanID, err1 = rt.db.NewBan(id, b.UserIDBanned, b.TimeStamp)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

// Removes a ban
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	id, err := strconv.Atoi(ps.ByName("userID"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

// struttura del ban
type ban struct {
	BanID        int    `json:"banID"`        // id del ban
	UserIDBanned int    `json:"userIDBanned"` // id dello user bannato
	TimeStamp    string `json:"timestamp"`    // timestamp di quando è avvenuto il follow
}
