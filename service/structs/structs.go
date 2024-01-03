package structs

// struttura dell'utente
type User struct {
	Id       int    `json:"userID"`   // id utente
	Username string `json:"username"` // username utente
}

// struttura del follow
type Follow struct {
	FollowId  int    `json:"followID"`  // id del follow
	UserID    int    `json:"userID"`    // id dello user che segue
	Followed  int    `json:"followed"`  // id dello user seguito
	TimeStamp string `json:"timestamp"` // timestamp di quando è avvenuto il follow
}

// struttura della foto
type Photo struct {
	PhotoID   int    `json:"photoID"`   // id della foto
	UserID    int    `json:"userID"`    // id dell'utente proprietario della foto
	Binary    []byte `json:"binary"`    // codifica binaria della foto
	TimeStamp string `json:"timestamp"` // timestamp di quando è stata postata la foto
}

// struttura del ban
type Ban struct {
	BanID     int    `json:"banID"`     // id del ban
	UserID    int    `json:"userID"`    // id dello user che banna
	Banned    int    `json:"banned"`    // id dello user bannato
	TimeStamp string `json:"timestamp"` // timestamp di quando è avvenuto il follow
}

// struttura del commento
type Comment struct {
	CommentID int    `json:"commentID"` // id del commento
	PhotoID   int    `json:"photoID"`   // id della foto
	Text      string `json:"text"`      // testo del commento
	User      User   `json:"user"`      // owner del commento
	TimeStamp string `json:"timestamp"` // timestamp di quando è stato postato il commento
}

// struttura del like
type Like struct {
	LikeID    int    `json:"likeID"`    // id del like
	PhotoID   int    `json:"photoID"`   // id della foto
	UserID    int    `json:"userID"`    // owner del commento
	TimeStamp string `json:"timestamp"` // timestamp di quando è stato postato il commento
}

// Profilo completo di un utente (Informazioni)
type Profile struct {
	User       User        `json:"user"`       // informazioni relative all'utente, ossia: userID, username
	Photos     []PhotoInfo `json:"photos"`     // foto dell'utente
	Followers  []User      `json:"followers"`  // seguaci dell'utente
	Followings []User      `json:"followings"` // seguiti dell'utente
	NPhotos    int         `json:"nPhotos"`    // numero di foto caricate dall'utente
	Bans       []User      `json:"bans"`       // ban dell'utente
}

// Insieme di foto dei following dell'utente
type Stream struct {
	User   User        `json:"user"`   // informazioni relative all'utente ossia: userID, username
	Photos []PhotoInfo `json:"photos"` // foto dello stream dell'utente
}

// Informazioni complete riguardanti la foto
type PhotoInfo struct {
	Photo     Photo     `json:"photo"`     // informazioni relative alla foto, ossia: photoID, userID, path e timestamp
	Likes     []Like    `json:"likes"`     // Like della foto
	Comments  []Comment `json:"comments"`  // Commenti della foto
	NLikes    int       `json:"nLikes"`    // Numero di likes della foto
	NComments int       `json:"nComments"` // Numero di commenti della foto

}
