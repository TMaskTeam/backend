package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/dto"
	"backend/src/internal/service/abstract"
	"errors"
	"net/http"
	"strconv"
)

func CreateBonusProgram(
	ctx context.HandlerContext,
	req *dto.BonusProgramRequest,
	bonusProgramService abstract.IBonusProgramService,
) (interface{}, error) {

	businessID, err := strconv.Atoi(ctx.Param("business_id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, errors.New("invalid business_id")
	}

	ownerID, ok := ctx.GetLocal("user_id").(int)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	program, err := bonusProgramService.Create(businessID, ownerID, req.ProgramName, req.TokenName)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}

	resp := dto.BonusProgramResponse{
		ProgramID:   program.ProgramID,
		BusinessID:  program.BusinessID,
		ProgramName: program.ProgramName,
		TokenName:   program.TokenName,
	}

	ctx.Status(http.StatusCreated)
	return resp, nil
}

func GetBonusProgramsByBusinessID(
	ctx context.HandlerContext,
	bonusProgramService abstract.IBonusProgramService,
) (interface{}, error) {
	businessID, err := strconv.Atoi(ctx.Get("business_id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}

	programs, err := bonusProgramService.GetByBusinessID(businessID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}

	resp := make([]dto.BonusProgramResponse, 0, len(programs))
	for _, p := range programs {
		resp = append(resp, dto.BonusProgramResponse{
			ProgramID:   p.ProgramID,
			BusinessID:  p.BusinessID,
			ProgramName: p.ProgramName,
			TokenName:   p.TokenName,
		})
	}

	ctx.Status(http.StatusOK)
	return resp, nil
}

func GetAllBonusPrograms(
	ctx context.HandlerContext,
	bonusProgramService abstract.IBonusProgramService,
) (interface{}, error) {
	programs, err := bonusProgramService.GetAll()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}

	resp := make([]dto.BonusProgramResponse, 0, len(programs))
	for _, p := range programs {
		resp = append(resp, dto.BonusProgramResponse{
			ProgramID:   p.ProgramID,
			BusinessID:  p.BusinessID,
			ProgramName: p.ProgramName,
			TokenName:   p.TokenName,
		})
	}

	ctx.Status(http.StatusOK)
	return resp, nil
}
