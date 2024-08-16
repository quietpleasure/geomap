package router

import "github.com/gin-gonic/gin"

func (r *router) publicRoutes(g *gin.RouterGroup) {

	g.GET("/tracks", r.handlers.AllTracks())
	g.POST("/tracks", r.handlers.ShowTracks())
	g.POST("/track", r.handlers.ShowTrack())
}

func (r *router) privateRoutes(g *gin.RouterGroup) {
	// g.GET("index", handlers.IndexGetHandler())
	// g.GET("logout", handlers.LogoutGetHandler())
	// g.POST("mark-message-read", handlers.MarkMessageRead())
}

func (r *router) adminRoutes(g *gin.RouterGroup) {
	// g.GET("users",  handlers.UsersHandler())
	// g.POST("users", handlers.UsersHandler())
	// g.POST("switch-user", handlers.SwitchUserActivity())
}
