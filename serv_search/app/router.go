package app

import (
	"github.com/gin-gonic/gin"
)

// MapUrls maps the urls
func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	router.GET("/search=:searchQuery", dependencies.SearchController.Search)
	router.GET("/search/byuser=:userid", dependencies.SearchController.SearchByUserId)
}
