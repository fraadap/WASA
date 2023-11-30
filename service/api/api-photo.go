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

// Posts a new photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	var ph structs.Photo
	err0 := json.Unmarshal(body, &ph)

	if ph.TimeStamp == "" {
		ph.TimeStamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	} else if len(ph.TimeStamp) != 20 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err0 != nil || ph.Path == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var err1 error
	ph.PhotoID, err1 = rt.db.NewPhoto(id, ph.Path, ph.TimeStamp)
	ph.UserID = id
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w).Encode(ph)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	photoId, err0 := strconv.Atoi(ps.ByName("photoID"))

	if err0 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err1 := rt.db.DeletePhoto(id, photoId)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	st, err := rt.db.GetMyStream(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(st)
}
