package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getUserProfile is
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "text/plain")
	_ = json.NewEncoder(w).Encode(id)

}

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var u user
	err0 := json.Unmarshal(body, &u)

	if err0 != nil || u.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err1 := rt.db.SetUsername(id, u.Username)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.Id = id

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//
}

// struttura dell'utente
type user struct {
	Id         int      `json:"userID"`     // id utente
	Username   string   `json:"username"`   // username utente
	Photos     []photo  `json:"photos"`     // array di foto dell'utente
	Followings []follow `json:"followings"` // array di seguiti dell'utente
	Bans       []ban    `json:"bans"`       // array di ban	dell'utente
}
