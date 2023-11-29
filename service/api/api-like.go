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

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID, err := strconv.Atoi(ps.ByName("userID"))
	photoID, err1 := strconv.Atoi(ps.ByName("photoID"))
	if err != nil || err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// lettura del body
	body, er := io.ReadAll(r.Body)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var li structs.Like

	// conversione del body in struct like
	err0 := json.Unmarshal(body, &li)

	if err0 != nil || li.UserID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	li.PhotoID = photoID

	token := getToken(r.Header.Get("Authorization"))
	if li.UserID != token {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// generazione timestamp se assente altimenti controllo del formato
	if li.TimeStamp == "" {
		li.TimeStamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	} else if len(li.TimeStamp) != 20 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// controllo se l'utente è bannato dalla persona proprietaria della photo
	userIDPhoto, err := rt.db.UserIDByPhoto(photoID)
	if err != nil || userID != userIDPhoto {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if banned, er1 := rt.db.IsBanned(userIDPhoto, li.UserID); banned {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if er1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var err2 error

	// aggiunta del commento al db
	li.LikeID, err2 = rt.db.NewLike(li.UserID, li.PhotoID, li.TimeStamp)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// risposta
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(li) // return: struttura like
}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//ricezione dei params userID, photoID e likeID con relativa gestione degli errori di conversione

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

	likeId, err1 := strconv.Atoi(ps.ByName("likeID"))
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := getToken(r.Header.Get("Authorization"))

	owner, err := rt.db.GetOwnerFromLikeID(likeId)
	if err != nil || owner != token {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// eliminazione del like
	err2 := rt.db.DeleteLike(likeId, photoId, id)
	if err2 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//risposta 204
	w.WriteHeader(http.StatusNoContent)
}
