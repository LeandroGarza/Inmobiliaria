package app

import (
	clients "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/clients/rabbitmq"
	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/config"
	controllers "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/controllers"
	service "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services"
	repositories "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services/repositories"
)

type Dependencies struct {
	ItemController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	ccache := repositories.NewCCache(config.CCMAXSIZE, config.CCITEMSTOPRUNE, config.CCDEFAULTTTL)
	mongo := repositories.NewMongoDB(config.DBHOST, config.DBPORT, config.COLLECTION)
	queue := clients.NewRabbitmq(config.RABBITHOST, config.RABBITPORT)

	// Services
	service := service.NewServiceImpl(ccache, mongo, queue)

	// consumer
	go queue.ConsumeUserUpdate(config.EXCHANGE, ccache, mongo)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}
