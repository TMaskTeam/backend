package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/dto"
	"backend/src/internal/service/abstract"
	"errors"
	"net/http"
)

func GetMe(
	ctx context.HandlerContext,
	ownerService abstract.IBusinessOwnerService,
	clientService abstract.IClientService,
) (interface{}, error) {
	userID, ok := ctx.GetLocal("user_id").(int)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	role, ok := ctx.GetLocal("role").(string)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	switch role {
	case "client":
		client, err := clientService.GetByID(userID)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return nil, err
		}

		resp := buildClientMeResponse(role, client)

		ctx.Status(http.StatusOK)
		return resp, nil

	case "business_owner":
		owner, err := ownerService.GetByID(userID)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return nil, errors.New("owner not found")
		}

		resp := buildOwnerMeResponse(role, owner)

		ctx.Status(http.StatusOK)
		return resp, nil

	default:
		ctx.Status(http.StatusBadRequest)
		return nil, errors.New("unknown role")
	}

}

func Update(
	ctx context.HandlerContext,
	ownerService abstract.IBusinessOwnerService,
	clientService abstract.IClientService,
) (interface{}, error) {
	userID, ok := ctx.GetLocal("user_id").(int)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	role, ok := ctx.GetLocal("role").(string)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	switch role {
	case "client":
		client, err := clientService.GetByID(userID)
		if err != nil {
			return nil, err
		}

		client.ID = userID
		clientNew, err := clientService.Update(client)
		if err != nil {
			return nil, err
		}

		resp := buildClientMeResponse(role, clientNew)

		ctx.Status(http.StatusOK)
		return resp, nil

	case "business_owner":
		owner, err := ownerService.GetByID(userID)
		if err != nil {
			return nil, err
		}

		owner.ID = userID
		ownerNew, err := ownerService.Update(owner)
		if err != nil {
			return nil, err
		}

		resp := buildOwnerMeResponse(role, ownerNew)

		ctx.Status(http.StatusOK)
		return resp, nil

	default:
		ctx.Status(http.StatusBadRequest)
		return nil, errors.New("unknown role")
	}

}

func buildClientMeResponse(role string, client *domain.Client) dto.ClientMeResponse {
	return dto.ClientMeResponse{
		Role:        role,
		ID:          client.ID,
		FirstName:   client.FirstName,
		MiddleName:  client.MiddleName,
		LastName:    client.LastName,
		PhoneNumber: client.PhoneNumber,
		Email:       client.Email,
		Birthday:    client.Birthday,
	}
}

func buildOwnerMeResponse(role string, owner *domain.BusinessOwner) dto.BusinessOwnerMeResponse {
	return dto.BusinessOwnerMeResponse{
		Role:        role,
		ID:          owner.ID,
		FirstName:   owner.FirstName,
		MiddleName:  owner.MiddleName,
		LastName:    owner.LastName,
		INN:         owner.INN,
		PhoneNumber: owner.PhoneNumber,
		Email:       owner.Email,
		Birthday:    owner.Birthday,
	}
}
