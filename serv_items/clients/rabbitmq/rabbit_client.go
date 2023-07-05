package rabbitmq

import (
	"context"

	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services/repositories"
)

type QueueClient interface {
	PublishItem(ctx context.Context, item dtos.ItemDto) error
	PublishItems(ctx context.Context, items dtos.ItemsDto) error
	ConsumeUserUpdate(exchange string, ccache repositories.Repository, mongo repositories.Repository)
}
