package dto

type ClientJoinProgramRequest struct {
	ProgramID int `json:"program_id"`
}

type ClientJoinProgramResponse struct {
	ClientBonusProgramID int `json:"client_bonus_program_id"`
}
