package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
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

	if id != token || u.Id != token {
		w.WriteHeader(http.StatusForbidden)
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

func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	text := ps.ByName("text")

	token := getToken(r.Header.Get("Authorization"))

	users, err := rt.db.SearchUsers(text)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Print(err.Error())
		return
	}

	var profiles []structs.Profile

	for _, u := range users {
		if u.Id == token {
			continue
		}
		p, err := rt.db.GetProfile(u.Id)

		if err != nil {
			fmt.Print(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p.Bans = nil
		p.Photos = nil
		profiles = append(profiles, p)

	}
	sort.Slice(profiles, func(i, j int) bool {
		return len(profiles[i].Followers) > len(profiles[j].Followers)
	})

	// manca il sort per numero di followers
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w).Encode(profiles)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
