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
	request *dto.ClientLoginRequest,
	clientService abstract.IClientService,
) (dto.ClientLoginResponse, error) {
	token, expiresAt, owner, err := clientService.Login(request.Login, request.Password)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.ClientLoginResponse{}, err
	}

	resp := buildClientLoginResponse(token, expiresAt, owner)

	ctx.Status(http.StatusOK)
	return resp, nil
}

func buildClientLoginResponse(token string, expiresAt time.Time, client *domain.Client) dto.ClientLoginResponse {
	return dto.ClientLoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		Owner: dto.ClientResponse{
			ID:          client.ID,
			FirstName:   client.FirstName,
			MiddleName:  client.MiddleName,
			LastName:    client.LastName,
			PhoneNumber: client.PhoneNumber,
			Email:       client.Email,
			Birthday:    client.Birthday,
		},
	}
}

func ClientJoin(
	ctx context.HandlerContext,
	request *dto.ClientJoinRequest,
	clientService abstract.IClientService,
) (dto.ClientJoinResponse, error) {
	client, err := model.ToDomain[dto.ClientResponse, domain.Client](&request.Client)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.ClientJoinResponse{}, err
	}

	clientBonusProgramID, TokensCount, err := clientService.Join(client, request.ProgramID)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.ClientJoinResponse{}, err
	}

	resp := buildClientJoinResponse(clientBonusProgramID, request.ProgramID, TokensCount, client)

	ctx.Status(http.StatusOK)
	return resp, nil
}

func buildClientJoinResponse(clientBonusProgramID, programID, tokensCount int, client *domain.Client) dto.ClientJoinResponse {
	return dto.ClientJoinResponse{
		ProgramID:            programID,
		ClientBonusProgramID: clientBonusProgramID,
		TokensCount:          tokensCount,
		Client: dto.ClientResponse{
			ID:          client.ID,
			FirstName:   client.FirstName,
			MiddleName:  client.MiddleName,
			LastName:    client.LastName,
			PhoneNumber: client.PhoneNumber,
			Email:       client.Email,
			Birthday:    client.Birthday,
		},
	}
}
