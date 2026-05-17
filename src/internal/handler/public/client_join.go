package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/dto"
	"backend/src/internal/model"
	"backend/src/internal/service/abstract"
	"net/http"
)

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

	clientBonusProgramID, tokensCount, err := clientService.Join(client, request.ProgramID)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.ClientJoinResponse{}, err
	}

	resp := buildClientJoinResponse(clientBonusProgramID, request.ProgramID, tokensCount, client)

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
