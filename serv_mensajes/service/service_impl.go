package service

import (
	"messages/dto"
	"messages/model"
	"messages/service/repositories"
	"messages/utils/errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type MessageServiceImpl struct {
	MessageClient repositories.MessageClient
}

func NewMessageServiceImpl(messageclient repositories.MessageClient) *MessageServiceImpl {
	messageclient.StartDbEngine()
	return &MessageServiceImpl{
		MessageClient: messageclient,
	}
}

func (s *MessageServiceImpl) CreateMessage(messagedto dto.MessageDto) (dto.MessageDto, errors.ApiError) {
	var message model.Message

	message.Userid = messagedto.Userid
	message.Itemid = messagedto.Itemid
	message.Content = messagedto.Content
	message.Createdat = time.Now().Format("2006-01-02 15:04:05")

	message, err := s.MessageClient.CreateMessage(message)
	if err != nil {
		return dto.MessageDto{}, errors.NewInternalServerApiError("Error creating new message", err)
	}
	messagedto.Id = message.Id

	return messagedto, nil
}

func (s *MessageServiceImpl) GetMessagesByItem(itemid string) (dto.MessagesDto, errors.ApiError) {
	var messagesdto dto.MessagesDto

	messages, err := s.MessageClient.GetMessagesByItem(itemid)
	if err != nil {
		return dto.MessagesDto{}, errors.NewInternalServerApiError("Error getting messages by item", err)
	}

	for _, message := range messages {
		var messagedto dto.MessageDto

		messagedto.Id = message.Id
		messagedto.Userid = message.Userid
		messagedto.Itemid = message.Itemid
		messagedto.Content = message.Content
		messagedto.Createdat = message.Createdat

		messagesdto = append(messagesdto, messagedto)
	}

	return messagesdto, nil
}

func (s *MessageServiceImpl) GetMessageById(id int) (dto.MessageDto, errors.ApiError) {
	var messagedto dto.MessageDto

	message, err := s.MessageClient.GetMessageById(id)
	if err != nil {
		return dto.MessageDto{}, errors.NewBadRequestApiError("Failed to get message by id")
	}

	messagedto.Id = message.Id
	messagedto.Userid = message.Userid
	messagedto.Itemid = message.Itemid
	messagedto.Content = message.Content
	messagedto.Createdat = message.Createdat

	return messagedto, nil
}

func (s *MessageServiceImpl) GetMessagesByUser(userid int) (dto.MessagesDto, errors.ApiError) {
	var messagesdto dto.MessagesDto

	messages, err := s.MessageClient.GetMessagesByUser(userid)
	if err != nil {
		return dto.MessagesDto{}, errors.NewBadRequestApiError("failed to get messages by user")
	}

	for _, message := range messages {
		var messagedto dto.MessageDto
		messagedto.Id = message.Id
		messagedto.Userid = message.Userid
		messagedto.Itemid = message.Itemid
		messagedto.Content = message.Content
		messagedto.Createdat = message.Createdat

		messagesdto = append(messagesdto, messagedto)
	}

	return messagesdto, nil
}

func (s *MessageServiceImpl) ValidateToken(tokenString string) (*jwt.Token, error) {
	// Definir la clave secreta utilizada para firmar el token
	var jwtKey = []byte("tengohambre")

	// Configurar el validador del token
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar el algoritmo de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewUnauthorizedApiError("Invalid token")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *MessageServiceImpl) ValidateRequest(r *http.Request) (*dto.Claims, error) {
	// Obtener el token del encabezado de la solicitud
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.NewBadRequestApiError("Missing authorization header")
	}

	// Extraer el token de autenticación del encabezado
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Validar el token
	token, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Obtener los claims del token
	claims, ok := token.Claims.(*dto.Claims)
	if !ok {
		return nil, errors.NewUnauthorizedApiError("Invalid token claims")
	}

	return claims, nil
}
