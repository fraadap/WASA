package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)

	// users section
	rt.router.POST("/session", rt.doLogin)             // 1)fatto, manca la parte di sicurezza con variabile di sessione
	rt.router.GET("/users/:userID", rt.getUserProfile) // 13)
	rt.router.PUT("/users/:userID", rt.setMyUsername)  // 2)fatto

	// followers section
	rt.router.POST("/users/:userID/followers", rt.followUser)               // 3)fatto
	rt.router.DELETE("/users/:userID/followers/:followID", rt.unfollowUser) // 4)fatto

	// photos section
	rt.router.GET("/users/:userID/photos", rt.getMyStream)             // 14)
	rt.router.POST("/users/:userID/photos", rt.uploadPhoto)            // 7)fatto
	rt.router.DELETE("/users/:userID/photos/:photoID", rt.deletePhoto) // 8)fatto

	// comments section
	rt.router.POST("/users/:userID/photos/:photoID/comments", rt.commentPhoto)                // 9)fatto
	rt.router.DELETE("/users/:userID/photos/:photoID/comments/:commentID", rt.uncommentPhoto) // 10)fatto

	// likes section
	rt.router.POST("/users/:userID/photos/:photoID/likes", rt.likePhoto)             // 11)fatto
	rt.router.DELETE("/users/:userID/photos/:photoID/likes/:likeID", rt.unlikePhoto) // 12)fatto

	// bans section
	rt.router.POST("/users/:userID/bans", rt.banUser)            // 5) fatto
	rt.router.DELETE("/users/:userID/bans/:banID", rt.unbanUser) // 6) fatto

	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
