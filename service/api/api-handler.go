package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)

	// users section
	rt.router.POST("/session", rt.wrap(rt.doLogin))                       // 1)fatto, manca la parte di sicurezza con variabile di sessione
	rt.router.GET("/users/:userID", rt.wrap(rt.getUserProfile))           // 13)fatto
	rt.router.PUT("/users/:userID", rt.wrap(rt.setMyUsername))            // 2)fatto
	rt.router.GET("/users/:userID/username", rt.wrap(rt.getUsernameByID)) // riportato nell'api

	rt.router.GET("/search/:text", rt.wrap(rt.searchUsers)) // riportato nell'api

	// followers section
	rt.router.PUT("/users/:userID/follow", rt.wrap(rt.followUser))                // 3)fatto
	rt.router.DELETE("/users/:userID/follow/:followID", rt.wrap(rt.unfollowUser)) // 4)fatto
	rt.router.GET("/users/:userID/followID/:followed", rt.wrap(rt.getFollowID))   // riportato nell'api

	// photos section
	rt.router.GET("/users/:userID/photos", rt.wrap(rt.getMyStream))             // 14)fatto
	rt.router.POST("/users/:userID/photos", rt.wrap(rt.uploadPhoto))            // 7)fatto
	rt.router.DELETE("/users/:userID/photos/:photoID", rt.wrap(rt.deletePhoto)) // 8)fatto

	// comments section
	rt.router.POST("/users/:userID/photos/:photoID/comments", rt.wrap(rt.commentPhoto))                // 9)fatto
	rt.router.DELETE("/users/:userID/photos/:photoID/comments/:commentID", rt.wrap(rt.uncommentPhoto)) // 10)fatto

	// likes section
	rt.router.PUT("/users/:userID/photos/:photoID/likes", rt.wrap(rt.likePhoto))              // 11)fatto
	rt.router.DELETE("/users/:userID/photos/:photoID/likes/:likeID", rt.wrap(rt.unlikePhoto)) // 12)fatto

	// bans section
	rt.router.PUT("/users/:userID/bans", rt.wrap(rt.banUser))             // 5) fatto
	rt.router.DELETE("/users/:userID/bans/:banID", rt.wrap(rt.unbanUser)) // 6) fatto
	rt.router.GET("/users/:userID/banID/:banned", rt.wrap(rt.getBanID))   // riportato nell'api

	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
