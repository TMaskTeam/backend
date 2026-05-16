package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/dto"
	"backend/src/internal/model"
	"backend/src/internal/service/abstract"
	"errors"
	"net/http"
	"time"
)

func RegisterClient(
	ctx context.HandlerContext,
	request *dto.ClientRegisterRequest,
	clientService abstract.IClientService,
) (dto.ClientRegisterResponse, error) {

	birthday, err := time.Parse("2006-01-02", request.Birthday)
	if err != nil {
		ctx.Status(400)
		return dto.ClientRegisterResponse{}, errors.New("invalid birthday format, expected YYYY-MM-DD")
	}

	client, err := model.ToDomain[dto.ClientRegisterRequest, domain.Client](request)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.ClientRegisterResponse{}, err
	}
	client.Birthday = birthday

	err = clientService.Register(client)
	if err != nil {
		ctx.Status(http.StatusConflict)
		return dto.ClientRegisterResponse{}, err
	}
	resp := buildClientResponse(client)

	ctx.Status(201)
	return resp, nil
}

func buildClientResponse(createdClient *domain.Client) dto.ClientRegisterResponse {
	return dto.ClientRegisterResponse{
		ID:          createdClient.ID,
		FirstName:   createdClient.FirstName,
		MiddleName:  createdClient.MiddleName,
		LastName:    createdClient.LastName,
		PhoneNumber: createdClient.PhoneNumber,
		Email:       createdClient.Email,
		Birthday:    createdClient.Birthday,
	}
}
