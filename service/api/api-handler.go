package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)

	//users section
	rt.router.POST("/session", rt.doLogin)
	rt.router.GET("/users/:userID", rt.getUserProfile)
	rt.router.PUT("/users/:userID", rt.setMyUsername)

	//followers section
	rt.router.POST("/users/:userID/followers", rt.followUser)
	rt.router.DELETE("/users/:userID/followers/:followID", rt.unfollowUser)

	//photos section
	rt.router.GET("/users/:userID/photos", rt.getMyStream)
	rt.router.POST("/users/:userID/photos", rt.uploadPhoto)
	rt.router.DELETE("/users/:userID/photos/:photoID", rt.deletePhoto)

	//comments section
	rt.router.POST("/users/:userID/photos/:photoID/comments", rt.commentPhoto)
	rt.router.DELETE("/users/:userID/photos/:photoID/comments/:commentID", rt.uncommentPhoto)

	//likes section
	rt.router.POST("/users/:userID/photos/:photoID/likes", rt.likePhoto)
	rt.router.DELETE("/users/:userID/photos/:photoID/likes/:likeID", rt.unlikePhoto)

	//bans section
	rt.router.POST("/users/:userID/bans", rt.banUser)
	rt.router.DELETE("/users/:userID/bans", rt.unbanUser)

	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
