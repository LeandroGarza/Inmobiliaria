package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	//log "github.com/sirupsen/logrus"

	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/clients/rabbitmq"
	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services/repositories"
	e "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/utils/errors"
	"github.com/golang-jwt/jwt"
)

type ServiceImpl struct {
	localCache repositories.Repository
	db         repositories.Repository
	queue      rabbitmq.QueueClient
}

func NewServiceImpl(localCache repositories.Repository, db repositories.Repository, queue rabbitmq.QueueClient) *ServiceImpl {
	return &ServiceImpl{
		localCache: localCache,
		db:         db,
		queue:      queue,
	}
}

func (serv *ServiceImpl) Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	var item dtos.ItemDto
	var apiErr e.ApiError
	var source string

	// try to find it in localCache
	item, apiErr = serv.localCache.Get(ctx, id)
	if apiErr != nil {
		if apiErr.Status() != http.StatusNotFound {
			return dtos.ItemDto{}, apiErr
		}
		// try to find it in db
		item, apiErr = serv.db.Get(ctx, id)
		if apiErr != nil {
			if apiErr.Status() != http.StatusNotFound {
				return dtos.ItemDto{}, apiErr
			} else {
				fmt.Printf("Not found item %s in any datasource", id)
				apiErr = e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
				return dtos.ItemDto{}, apiErr
			}
		} else {
			source = "db"
			defer func() {
				if _, apiErr := serv.localCache.InsertItem(ctx, item); apiErr != nil {
					fmt.Printf("Error trying to save item in localCache %v", apiErr)
				}
			}()
		}
	} else {
		source = "localCache"
	}

	fmt.Printf("Obtained item from %s!", source)
	return item, nil
}

func downloadImage(url string, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return e.NewNotFoundApiError("image not found")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func (serv *ServiceImpl) InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	result, apiErr := serv.db.InsertItem(ctx, item)
	if apiErr != nil {
		fmt.Println(fmt.Printf("Error inserting item in db: %v\n", apiErr))
		return dtos.ItemDto{}, apiErr
	}
	fmt.Println(fmt.Printf("Inserted item in db: %v\n", result))

	if err := serv.queue.PublishItem(ctx, result); err != nil {
		return result, e.NewInternalServerApiError(fmt.Sprintf("Error publishing item %s", item.Id), err)
	}
	fmt.Println(fmt.Printf("Message sent: %v\n", result.Id))

	go downloadImage(result.Image, result.Id)

	return result, nil
}

func (serv *ServiceImpl) InsertItems(ctx context.Context, items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError) {
	result, apiErr := serv.db.InsertItems(ctx, items)
	if apiErr != nil {
		fmt.Println(fmt.Printf("Error inserting items in db: %v\n", apiErr))
		return dtos.ItemsDto{}, apiErr
	}
	fmt.Println(fmt.Printf("Inserted items in db: %v\n", result))

	if err := serv.queue.PublishItems(ctx, result); err != nil {
		return result, e.NewInternalServerApiError("Error publishing items", err)
	}
	fmt.Println(fmt.Printf("Message sent"))

	for _, item := range result {
		go downloadImage(item.Image, item.Id)
	}
	return result, nil
}

func (v *ServiceImpl) ValidateToken(tokenString string) (*jwt.Token, error) {
	// Definir la clave secreta utilizada para firmar el token
	var jwtKey = []byte("tengohambre")

	// Configurar el validador del token
	token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar el algoritmo de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (v *ServiceImpl) ValidateRequest(r *http.Request) (*dtos.Claims, error) {
	// Obtener el token del encabezado de la solicitud
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("missing authorization header")
	}

	// Extraer el token de autenticación del encabezado
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Validar el token
	token, err := v.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Obtener los claims del token
	claims, ok := token.Claims.(*dtos.Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
