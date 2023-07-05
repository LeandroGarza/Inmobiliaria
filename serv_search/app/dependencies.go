package app

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/config"
	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/controllers"
	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/services"
	repositories "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/services/repositories"
)

type Dependencies struct {
	SearchController *controllers.SearchController
}

func BuildDependencies() *Dependencies {
	// repositories
	searchclient := repositories.NewSearchClient()
	queueclient := repositories.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, "localhost", config.RABBITPORT)

	// services
	service := services.NewSearchService(searchclient, queueclient)

	// controllers
	searchcontroller := controllers.NewSearchController(service)

	// consumers
	go queueclient.ConsumeItems()
	go queueclient.ConsumeUserUpdates(config.EXCHANGE, searchclient)

	return &Dependencies{
		SearchController: searchcontroller,
	}
}
