package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Posts a new photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	var ph photo
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
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ph)
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("userID"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

// struttura della foto
type photo struct {
	PhotoID   int       `json:"photoID"`   // id della foto
	Path      string    `json:"path"`      // path della foto
	TimeStamp string    `json:"timestamp"` // timestamp di quando è stata postata la foto
	Comments  []comment `json:"comments"`  // array di commenti della foto
	Likes     []like    `json:"likes"`     // array di like della foto
}
