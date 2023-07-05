package service

import (
	"messages/dto"
	"messages/utils/errors"
	"net/http"
)

type MessageService interface {
	CreateMessage(messagedto dto.MessageDto) (dto.MessageDto, errors.ApiError)
	GetMessagesByItem(itemid string) (dto.MessagesDto, errors.ApiError)
	GetMessageById(id int) (dto.MessageDto, errors.ApiError)
	GetMessagesByUser(userid int) (dto.MessagesDto, errors.ApiError)
	ValidateRequest(r *http.Request) (*dto.Claims, error)
}
