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

func ClientRegister(
	ctx context.HandlerContext,
	request *dto.ClientRegisterRequest,
	clientService abstract.IClientService,
) (dto.ClientResponse, error) {

	birthday, err := time.Parse("2006-01-02", request.Birthday)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.ClientResponse{}, errors.New("invalid birthday format, expected YYYY-MM-DD")
	}

	client, err := model.ToDomain[dto.ClientRegisterRequest, domain.Client](request)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.ClientResponse{}, err
	}
	client.Birthday = birthday

	err = clientService.Register(client)
	if err != nil {
		ctx.Status(http.StatusConflict)
		return dto.ClientResponse{}, err
	}
	resp := buildClientRegisterResponse(client)

	ctx.Status(http.StatusOK)
	return resp, nil
}

func buildClientRegisterResponse(createdClient *domain.Client) dto.ClientResponse {
	return dto.ClientResponse{
		ID:          createdClient.ID,
		FirstName:   createdClient.FirstName,
		MiddleName:  createdClient.MiddleName,
		LastName:    createdClient.LastName,
		PhoneNumber: createdClient.PhoneNumber,
		Email:       createdClient.Email,
		Birthday:    createdClient.Birthday,
	}
}

func ClientLogin(
	ctx context.HandlerContext,
	request *dto.BusinessOwnerLoginRequest,
	ownerService abstract.IBusinessOwnerService,
) (dto.BusinessOwnerLoginResponse, error) {
	token, expiresAt, owner, err := ownerService.Login(request.Login, request.Password)
	if err != nil {
		return dto.BusinessOwnerLoginResponse{}, err
	}

	resp := buildClientLoginResponse(token, expiresAt, owner)

	ctx.Status(http.StatusOK)
	return resp, nil
}

func buildClientLoginResponse(token string, expiresAt time.Time, owner *domain.BusinessOwner) dto.BusinessOwnerLoginResponse {
	return dto.BusinessOwnerLoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		Owner: dto.BusinessOwnerResponse{
			ID:          owner.ID,
			FirstName:   owner.FirstName,
			MiddleName:  owner.MiddleName,
			LastName:    owner.LastName,
			INN:         owner.INN,
			PhoneNumber: owner.PhoneNumber,
			Email:       owner.Email,
			Birthday:    owner.Birthday,
		},
	}
}
