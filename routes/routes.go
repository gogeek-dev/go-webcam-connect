package routes

import (
	"database/sql"
	controller "webcam/controllers"
	middlewareauth "webcam/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {

	router := gin.Default()

	rout := router.Group("")
	// start capturing
	go controller.Getframes()

	router.LoadHTMLGlob("templates/*/*.html")

	router.Static("/assets", "./assets")

	router.GET("/", controller.Loginview)

	router.GET("/newregister", controller.Register)

	router.POST("/savedata", controller.Regsave)

	router.POST("/signin", controller.Login)

	rout.Use(middlewareauth.Middleware())

	rout.GET("/index", controller.Index)

	rout.GET("/takepicture", controller.Takepicture)

	router.GET("/video", controller.Video)

	rout.GET("/savepicture", controller.Viewpicture)

	rout.GET("/saveimage", controller.Savepicture)

	rout.GET("/logout", controller.Logout)

	return router
}
