package services

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/dtos"
)

type Service interface {
	Search(query string) (dtos.ItemsDto, error)
	SearchByUserId(id int) (dtos.ItemsDto, error)
}
