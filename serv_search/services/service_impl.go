package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/dtos"
	clients "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/services/repositories"
)

type SearchService struct {
	searchclient *clients.SearchClient
	queueclient  *clients.QueueClient
}

func NewSearchService(searchclient *clients.SearchClient, queueclient *clients.QueueClient) *SearchService {
	return &SearchService{
		searchclient: searchclient,
		queueclient:  queueclient,
	}
}

func (s *SearchService) Search(query string) (dtos.ItemsDto, error) {
	// llamada al search de clients
	r, err := s.searchclient.Search(query)
	if err != nil {
		return dtos.ItemsDto{}, err
	}

	// lectura de la response y guardamos los bytes
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return dtos.ItemsDto{}, err
	}

	// llamada a la funcion que parsea items
	rdto, err := parseItems(bytes)
	if err != nil {
		return dtos.ItemsDto{}, err
	}

	return rdto.Response.Docs, nil
}

func (s *SearchService) SearchByUserId(id int) (dtos.ItemsDto, error) {
	r, err := s.searchclient.SearchByUserId(id)
	if err != nil {
		return dtos.ItemsDto{}, err
	}

	// lectura de la response y guardamos los bytes
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return dtos.ItemsDto{}, err
	}

	// llamada a la funcion que parsea items
	rdto, err := parseItems(bytes)
	if err != nil {
		return dtos.ItemsDto{}, err
	}

	return rdto.Response.Docs, nil
}

func parseItems(bytes []byte) (dtos.ResponseDto, error) {
	var rdto dtos.ResponseDto
	// decodificamos los bytes y los guardamos en la variable que
	if err := json.Unmarshal(bytes, &rdto); err != nil {
		return dtos.ResponseDto{}, err
	}
	// Imprimir el contenido de rdto para verificarlo
	fmt.Printf("Contenido de rdto: %+v\n", rdto)
	fmt.Printf("Contenido de Docs de rdto: %+v\n", rdto.Response.Docs)

	return rdto, nil
}
