package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
	"errors"
)

type BusinessService struct {
	conn         connection.IDBConnection
	businessRepo repository.IBusinessRepository
}

func NewBusinessService(
	conn connection.IDBConnection,
	businessRepo repository.IBusinessRepository,
) *BusinessService {
	return &BusinessService{
		conn:         conn,
		businessRepo: businessRepo,
	}
}

func (bs *BusinessService) Create(ownerID int, name string, address string) (*domain.Business, error) {
	business := &domain.Business{
		OwnerID: ownerID,
		Name:    name,
		Address: address,
	}

	err := bs.businessRepo.Create(bs.conn, business)
	if err != nil {
		return nil, err
	}

	return business, nil
}

func (bs *BusinessService) GetByOwnerID(ownerID int) ([]domain.Business, error) {
	businesses, err := bs.businessRepo.GetByOwnerID(bs.conn, ownerID)
	if err != nil {
		return nil, err
	}

	if businesses == nil {
		return []domain.Business{}, nil
	}

	return businesses, nil
}

func (bs *BusinessService) Delete(businessID, ownerID int) error {
	business, err := bs.businessRepo.GetByBusinessID(bs.conn, businessID)
	if err != nil {
		return err
	}
	if business == nil {
		return errors.New("business not found")
	}
	if business.OwnerID != ownerID {
		return errors.New("you don't have permission to delete this business")
	}

	return bs.businessRepo.Delete(bs.conn, businessID)
}
