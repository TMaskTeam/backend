package dto

type ClientJoinRequest struct {
	Client    ClientResponse `json:"client"`
	ProgramID int            `json:"program_id"`
}

type ClientJoinResponse struct {
	Client               ClientResponse `json:"client"`
	ClientBonusProgramID int            `json:"client_bonus_program_id"`
	ProgramID            int            `json:"program_id"`
	TokensCount          int            `json:"tokens_count"`
}
