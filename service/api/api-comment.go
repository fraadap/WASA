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

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	var com structs.Comment

	// conversione del body in struct comment
	err0 := json.Unmarshal(body, &com)

	if err0 != nil || com.Text == "" || com.User.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token := getToken(r.Header.Get("Authorization"))
	if com.User.Id != token {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// generazione timestamp se assente altimenti controllo del formato
	if com.TimeStamp == "" {
		com.TimeStamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	} else if len(com.TimeStamp) != 20 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	com.PhotoID = photoID

	userIDPhoto, err := rt.db.UserIDByPhoto(photoID)
	if err != nil || userID != userIDPhoto {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if banned, er1 := rt.db.IsBanned(userIDPhoto, com.User.Id); banned {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if er1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var err2 error

	// aggiunta del commento al db
	com.CommentID, err2 = rt.db.NewComment(com.User.Id, com.PhotoID, com.Text, com.TimeStamp)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	com.User.Username, err = rt.db.GetUsername(com.User.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// risposta
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w).Encode(com) // return: struttura comment
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// ricezione dei params userID, photoID e commentID con relativa gestione degli errori di conversione

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

	commentId, err1 := strconv.Atoi(ps.ByName("commentID"))
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := getToken(r.Header.Get("Authorization"))

	owner, err := rt.db.GetOwnerFromCommentID(commentId)
	if err != nil || owner != token {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// eliminazione del commento
	err2 := rt.db.DeleteComment(commentId, photoId, id)
	if err2 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// risposta 204
	w.WriteHeader(http.StatusNoContent)
}
