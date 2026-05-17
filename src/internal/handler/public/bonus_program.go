package public

import (
	context "backend/src/internal/context/abstract"
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

	resp := dto.ToBonusProgramResponse(program)

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

	resp := dto.ToBonusProgramResponseList(programs)

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

	resp := dto.ToBonusProgramResponseList(programs)

	ctx.Status(http.StatusOK)
	return resp, nil
}

func GetBonusProgramByProgramID(
	ctx context.HandlerContext,
	bonusProgramService abstract.IBonusProgramService,
) (*dto.BonusProgramResponse, error) {
	programID, err := strconv.Atoi(ctx.Params("program_id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}

	program, err := bonusProgramService.GetByProgramID(programID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}
	if program == nil {
		ctx.Status(http.StatusNotFound)
		return nil, nil
	}

	resp := dto.ToBonusProgramResponse(program)

	ctx.Status(http.StatusOK)
	return &resp, nil
}

func UpdateBonusProgram(
	ctx context.HandlerContext,
	bonusProgramService abstract.IBonusProgramService,
) (*dto.BonusProgramResponse, error) {
	programID, err := strconv.Atoi(ctx.Params("program_id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}

	var req dto.UpdateBonusProgramRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}

	program, err := bonusProgramService.Update(programID, req.ProgramName, req.TokenName)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}
	if program == nil {
		ctx.Status(http.StatusNotFound)
		return nil, nil
	}

	resp := dto.ToBonusProgramResponse(program)

	ctx.Status(http.StatusOK)
	return &resp, nil
}

func DeleteBonusProgram(
	ctx context.HandlerContext,
	bonusProgramService abstract.IBonusProgramService,
) (interface{}, error) {
	programID, err := strconv.Atoi(ctx.Params("program_id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}

	err = bonusProgramService.Delete(programID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil, err
	}

	ctx.Status(http.StatusNoContent)
	return nil, nil
}
