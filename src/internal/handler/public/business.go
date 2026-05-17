package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/dto"
	"backend/src/internal/service/abstract"
	"errors"
	"net/http"
	"strconv"
)

func CreateBusiness(
	ctx context.HandlerContext,
	req *dto.CreateBusinessRequest,
	businessService abstract.IBusinessService,
) (dto.CreateBusinessResponse, error) {

	ownerID, ok := ctx.GetLocal("user_id").(int)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return dto.CreateBusinessResponse{}, errors.New("unauthorized")
	}

	role, ok := ctx.GetLocal("role").(string)
	if !ok || role != "business_owner" {
		ctx.Status(http.StatusForbidden)
		return dto.CreateBusinessResponse{}, errors.New("only business owners can create businesses")
	}

	created, err := businessService.Create(ownerID, req.Name, req.Address)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return dto.CreateBusinessResponse{}, err
	}

	ctx.Status(http.StatusCreated)
	return buildCreateBusinessResponse(created), nil
}

func GetAllBusinesses(
	ctx context.HandlerContext,
	businessService abstract.IBusinessService,
) (dto.BusinessesListResponse, error) {

	ownerID, ok := ctx.GetLocal("user_id").(int)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return dto.BusinessesListResponse{}, errors.New("unauthorized")
	}

	role, ok := ctx.GetLocal("role").(string)
	if !ok || role != "business_owner" {
		ctx.Status(http.StatusForbidden)
		return dto.BusinessesListResponse{}, errors.New("only business owners can view businesses")
	}

	businesses, err := businessService.GetByOwnerID(ownerID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return dto.BusinessesListResponse{}, err
	}

	resp := buildBusinessesListResponse(businesses)

	ctx.Status(http.StatusOK)
	return resp, nil
}

func DeleteBusiness(
	ctx context.HandlerContext,
	businessService abstract.IBusinessService,
) (interface{}, error) {

	businessIDStr := ctx.Param("business_id")
	businessID, err := strconv.Atoi(businessIDStr)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, errors.New("invalid business_id")
	}

	ownerID, ok := ctx.GetLocal("user_id").(int)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	role, ok := ctx.GetLocal("role").(string)
	if !ok || role != "business_owner" {
		ctx.Status(http.StatusForbidden)
		return nil, errors.New("only business owners can delete businesses")
	}

	err = businessService.Delete(businessID, ownerID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}

	ctx.Status(http.StatusNoContent)
	return nil, nil
}

func buildBusinessesListResponse(businesses []domain.Business) dto.BusinessesListResponse {
	resp := dto.BusinessesListResponse{
		Businesses: make([]dto.BusinessResponse, 0, len(businesses)),
	}
	for _, b := range businesses {
		resp.Businesses = append(resp.Businesses, dto.BusinessResponse{
			BusinessID: b.BusinessID,
			Name:       b.Name,
			Address:    b.Address,
		})
	}
	return resp
}

func buildCreateBusinessResponse(created *domain.Business) dto.CreateBusinessResponse {
	return dto.CreateBusinessResponse{
		BusinessID: created.BusinessID,
		OwnerID:    created.OwnerID,
		Name:       created.Name,
		Address:    created.Address,
	}
}
