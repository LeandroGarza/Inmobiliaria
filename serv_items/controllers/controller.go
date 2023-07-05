package controller

import (
	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
	service "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services"
	e "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/utils/errors"

	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Get(c *gin.Context) {
	item, apiErr := ctrl.service.Get(c.Request.Context(), c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *Controller) InsertItem(c *gin.Context) {
	var item dtos.ItemDto
	if err := c.BindJSON(&item); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := ctrl.service.InsertItem(c.Request.Context(), item)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (ctrl *Controller) InsertItems(c *gin.Context) {
	var itemsdto dtos.ItemsDto
	if err := c.BindJSON(&itemsdto); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	/*userid := c.MustGet("userid").(int)

	for i := range itemsdto {
		itemsdto[i].UserId = userid
	}*/

	items, apirErr := ctrl.service.InsertItems(c.Request.Context(), itemsdto)
	if apirErr != nil {
		c.JSON(apirErr.Status(), apirErr)
		return
	}

	c.JSON(http.StatusCreated, items)
}

func (ctrl *Controller) ValidateTokenAndRequest(c *gin.Context) {

	claims, err := ctrl.service.ValidateRequest(c.Request)
	if err != nil {
		apiErr := e.NewUnauthorizedApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.Set("userid", claims.Userid)
	c.Set("username", claims.Username)
	c.Next()
}
