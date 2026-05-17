package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/dto"
	"backend/src/internal/service/abstract"
	"net/http"
	"strconv"
)

func CreateBonusProgram(
	ctx context.HandlerContext,
	bonusProgramService abstract.IBonusProgramService,
) (*dto.BonusProgramResponse, error) {
	businessID, err := strconv.Atoi(ctx.Params("business_id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}

	var req dto.BonusProgramRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}

	program, err := bonusProgramService.Create(businessID, req.ProgramName, req.TokenName)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}

	resp := buildBonusProgramResponse(program)

	ctx.Status(http.StatusCreated)
	return &resp, nil
}

func GetBonusProgramsByBusinessID(
	ctx context.HandlerContext,
	bonusProgramService abstract.IBonusProgramService,
) ([]dto.BonusProgramResponse, error) {
	businessID, err := strconv.Atoi(ctx.Params("business_id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}

	programs, err := bonusProgramService.GetByBusinessID(businessID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}

	resp := buildBonusProgramResponseList(programs)

	ctx.Status(http.StatusOK)
	return resp, nil
}

func GetAllBonusPrograms(
	ctx context.HandlerContext,
	bonusProgramService abstract.IBonusProgramService,
) ([]dto.BonusProgramResponse, error) {
	programs, err := bonusProgramService.GetAll()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}

	resp := buildBonusProgramResponseList(programs)

	ctx.Status(http.StatusOK)
	return resp, nil
}

func buildBonusProgramResponse(program *domain.BonusProgram) dto.BonusProgramResponse {
	return dto.BonusProgramResponse{
		ProgramID:   program.ProgramID,
		BusinessID:  program.BusinessID,
		ProgramName: program.ProgramName,
		TokenName:   program.TokenName,
	}
}

func buildBonusProgramResponseList(programs []*domain.BonusProgram) []dto.BonusProgramResponse {
	result := make([]dto.BonusProgramResponse, 0, len(programs))
	for _, p := range programs {
		result = append(result, buildBonusProgramResponse(p))
	}
	return result
}
