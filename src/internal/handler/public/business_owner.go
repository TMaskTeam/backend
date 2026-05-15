package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/dto"
	"backend/src/internal/model"
	"backend/src/internal/service/abstract"
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

	valid := validateOwner(owner)
}

func validateOwner(owner *domain.BusinessOwner) bool {

}
