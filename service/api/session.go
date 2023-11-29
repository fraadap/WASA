package api

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/fraadap/WASA/service/api/reqcontext"
	"github.com/fraadap/WASA/service/structs"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u structs.User
	err = json.Unmarshal(body, &u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := rt.db.Login(u.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(id)

}

func getToken(message string) int {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	stringToken := re.FindAllString(message, -1)
	token, _ := strconv.Atoi(stringToken[0])
	return token
}
