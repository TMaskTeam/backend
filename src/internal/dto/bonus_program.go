package dto

import "backend/src/internal/domain"

type BonusProgramRequest struct {
	ProgramName string `json:"program_name"`
	TokenName   string `json:"token_name"`
}

type BonusProgramResponse struct {
	ProgramID   int    `json:"program_id"`
	BusinessID  int    `json:"business_id"`
	ProgramName string `json:"program_name"`
	TokenName   string `json:"token_name"`
}

type UpdateBonusProgramRequest struct {
	ProgramName string `json:"program_name"`
	TokenName   string `json:"token_name"`
}

func ToBonusProgramResponse(program *domain.BonusProgram) BonusProgramResponse {
	return BonusProgramResponse{
		ProgramID:   program.ProgramID,
		BusinessID:  program.BusinessID,
		ProgramName: program.ProgramName,
		TokenName:   program.TokenName,
	}
}

func ToBonusProgramResponseList(programs []*domain.BonusProgram) []BonusProgramResponse {
	result := make([]BonusProgramResponse, 0, len(programs))
	for _, p := range programs {
		result = append(result, ToBonusProgramResponse(p))
	}
	return result
}
