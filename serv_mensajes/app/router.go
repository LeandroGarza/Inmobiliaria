package app

import "github.com/gin-gonic/gin"

func mapUrls(router *gin.Engine, dependencies *Dependencies) {
	router.GET("/messages/:id", dependencies.Controller.GetMessageById)
	router.GET("/messages/item/:itemid", dependencies.Controller.GetMessagesByItem)

	//router.Use(dependencies.Controller.ValidateRequestAndToken)
	router.POST("/messages", dependencies.Controller.CreateMessage)
	router.GET("/messages/user/:userid", dependencies.Controller.GetMessageByUser)
}
