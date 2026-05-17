package dto

type CreateBusinessRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=100"`
	Address string `json:"address" validate:"required"`
}

type CreateBusinessResponse struct {
	BusinessID int    `json:"business_id"`
	OwnerID    int    `json:"owner_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
}

type BusinessResponse struct {
	BusinessID int    `json:"business_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
}

type BusinessesListResponse struct {
	Businesses []BusinessResponse `json:"businesses"`
}
