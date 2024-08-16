package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//uuid "github.com/matoous/go-nanoid/v2"
)

func (r *router) gin(cfg *Config) *gin.Engine {
	gin.SetMode(cfg.GinMode)
	e := gin.New()
	e.Use(gin.Recovery())
	// e.SetFuncMap(template.FuncMap{
	// 	"CopyrightYear": func() int {
	// 		return time.Now().Year()
	// 	},
	// })
	// e.LoadHTMLGlob(cfg.HtmlTemplates + "/**/*")
	// e.Static("/"+cfg.StaticFiles, "./"+cfg.StaticFiles)
	// e.StaticFile("/favicon.ico", "./"+cfg.StaticFiles+"/images/favicon.ico")

	e.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{
			"Accept",
			"Accept-Language",
			"Origin",
			"Content-Language",
			"Access-Control-Allow-Origin",
			"Content-Type",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		//  return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	// secret, _ := uuid.New()
	// store := cookie.NewStore(secret)
	// store.Options(sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   int(10800),
	// 	HttpOnly: true,
	// 	SameSite: http.SameSiteLaxMode,
	// 	// Secure: true,
	// })

	// gob.Register(&account.User{})

	// e.Use(sessions.Sessions("user-session", store))

	middleware := r.middleware

	e.Use(middleware.ResponseRequestLogger)

	// e.Use(middleware.NotFound)

	public := e.Group("/api")
	r.publicRoutes(public)

	privat := e.Group("/")
	// private.Use(middleware.AuthRequiredAndAccesRights(account.ADMIN, account.USER))
	r.privateRoutes(privat)

	admin := e.Group("/")
	// admin.Use(middleware.AuthRequiredAndAccesRights(account.ADMIN))
	r.adminRoutes(admin)
	return e
}
