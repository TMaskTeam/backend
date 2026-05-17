package dto

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
