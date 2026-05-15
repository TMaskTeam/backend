package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/dto"
	"backend/src/internal/model"
	"backend/src/internal/service/abstract"
	"backend/src/internal/validator"
	"net/http"
)

func Register(
	ctx context.HandlerContext,
	request *dto.BusinessOwnerRegisterRequest,
	ownerService abstract.IBusinessOwnerService,
) (dto.BusinessOwnerRegisterResponse, error) {

	owner, err := model.ToDomain[dto.BusinessOwnerRegisterRequest, domain.BusinessOwner](request)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.BusinessOwnerRegisterResponse{}, err
	}

	err = validator.ValidateBusinessOwner(owner)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.BusinessOwnerRegisterResponse{}, err
	}

	err = ownerService.Register(owner)
	if err != nil {
		ctx.Status(http.StatusConflict)
		return dto.BusinessOwnerRegisterResponse{}, err
	}

	resp := buildResponse(owner)

	ctx.Status(201)
	return resp, nil
}

func buildResponse(createdOwner *domain.BusinessOwner) dto.BusinessOwnerRegisterResponse {
	return dto.BusinessOwnerRegisterResponse{
		ID:          createdOwner.ID,
		FirstName:   createdOwner.FirstName,
		MiddleName:  createdOwner.MiddleName,
		LastName:    createdOwner.LastName,
		INN:         createdOwner.INN,
		PhoneNumber: createdOwner.PhoneNumber,
		Email:       createdOwner.Email,
		Birthday:    createdOwner.Birthday,
	}
}
