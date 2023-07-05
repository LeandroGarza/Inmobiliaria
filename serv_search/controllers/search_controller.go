package controllers

import (
	"net/http"
	"strconv"

	service "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type SearchController struct {
	service service.Service
}

func NewSearchController(service service.Service) *SearchController {
	return &SearchController{
		service: service,
	}
}

func (sc *SearchController) Search(c *gin.Context) {
	query := c.Param("searchQuery")
	itemsDto, err := sc.service.Search(query)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)
}

func (sc *SearchController) SearchByUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	itemsdto, er := sc.service.SearchByUserId(id)
	if er != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, itemsdto)
}
