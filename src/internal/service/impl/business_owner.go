package impl

import (
	connection "backend/src/internal/db/abstract"
	repository "backend/src/internal/repository/abstract"
)

type BusinessOwnerService struct {
	conn      connection.IDBConnection
	ownerRepo repository.IBusinessOwnerRepository
}

func NewBusinessOwnerService(
	conn connection.IDBConnection,
	ownerRepo repository.IBusinessOwnerRepository,
) *BusinessOwnerService {
	return &BusinessOwnerService{
		conn:      conn,
		ownerRepo: ownerRepo,
	}
}
