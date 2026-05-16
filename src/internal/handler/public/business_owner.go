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

func Register(
	ctx context.HandlerContext,
	request *dto.BusinessOwnerRegisterRequest,
	ownerService abstract.IBusinessOwnerService,
) (dto.BusinessOwnerRegisterResponse, error) {

	birthday, err := time.Parse("2006-01-02", request.Birthday)
	if err != nil {
		ctx.Status(400)
		return dto.BusinessOwnerRegisterResponse{}, errors.New("invalid birthday format, expected YYYY-MM-DD")
	}

	owner, err := model.ToDomain[dto.BusinessOwnerRegisterRequest, domain.BusinessOwner](request)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.BusinessOwnerRegisterResponse{}, err
	}
	owner.Birthday = birthday

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
