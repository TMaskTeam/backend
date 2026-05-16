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
) (dto.BusinessOwnerResponse, error) {

	birthday, err := time.Parse("2006-01-02", request.Birthday)
	if err != nil {
		ctx.Status(400)
		return dto.BusinessOwnerResponse{}, errors.New("invalid birthday format, expected YYYY-MM-DD")
	}

	owner, err := model.ToDomain[dto.BusinessOwnerRegisterRequest, domain.BusinessOwner](request)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return dto.BusinessOwnerResponse{}, err
	}
	owner.Birthday = birthday

	err = ownerService.Register(owner)
	if err != nil {
		ctx.Status(http.StatusConflict)
		return dto.BusinessOwnerResponse{}, err
	}
	resp := buildRegisterResponse(owner)

	ctx.Status(201)
	return resp, nil
}

func buildRegisterResponse(createdOwner *domain.BusinessOwner) dto.BusinessOwnerResponse {
	return dto.BusinessOwnerResponse{
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

func Login(
	ctx context.HandlerContext,
	request *dto.BusinessOwnerLoginRequest,
	ownerService abstract.IBusinessOwnerService,
) (dto.BusinessOwnerLoginResponse, error) {

}
