package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var f follow
	err0 := json.Unmarshal(body, &f)

	if err0 != nil || f.FollowedId == 0 || f.FollowedId == id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// controllo se l'utente è bannato dalla persona che vuole seguire
	if banned, er1 := rt.db.IsBanned(f.FollowedId, id); banned {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if er1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if f.TimeStamp == "" {
		f.TimeStamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	} else if len(f.TimeStamp) != 20 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var err1 error
	f.FollowId, err1 = rt.db.NewFollow(id, f.FollowedId, f.TimeStamp)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(f)
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("userID"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	followId, err0 := strconv.Atoi(ps.ByName("followID"))

	if err0 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err1 := rt.db.DeleteFollow(id, followId)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// struttura del follow
type follow struct {
	FollowId   int    `json:"followID"`   // id del follow
	FollowedId int    `json:"followedID"` // id dello user seguito
	TimeStamp  string `json:"timestamp"`  // timestamp di quando è avvenuto il follow
}
