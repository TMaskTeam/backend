package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
)

type ClientBonusProgramRepository struct{}

func NewClientBonusProgramRepository() *ClientBonusProgramRepository {
	return &ClientBonusProgramRepository{}
}

func (cbp *ClientBonusProgramRepository) Upsert(conn abstract.IDBConnection, clientBonusProgram *domain.ClientBonusProgram) error {
	db := conn.Get().(*gorm.DB)

	clientBonusProgramDAO := &model.ClientBonusProgram{}
	clientBonusProgramDAO, err := clientBonusProgramDAO.ToModel(clientBonusProgram)
	if err != nil {
		return err
	}
	return db.Save(clientBonusProgramDAO).Error
}

func (cbp *ClientBonusProgramRepository) Delete(conn abstract.IDBConnection, bonusID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("client_bonus_program_id = ?", bonusID).Delete(&model.ClientBonusProgram{}).Error
}

func (cbp *ClientBonusProgramRepository) GetByClientID(conn abstract.IDBConnection, clientID int) (*domain.ClientBonusProgram, error) {
	db := conn.Get().(*gorm.DB)
	var clientBonusProgramDAO model.ClientBonusProgram
	err := db.Where("client_id = ?", clientID).First(&clientBonusProgramDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return clientBonusProgramDAO.ToDomain()
}

func (cbp *ClientBonusProgramRepository) GetByProgramID(conn abstract.IDBConnection, programID int) (*domain.ClientBonusProgram, error) {
	db := conn.Get().(*gorm.DB)
	var clientBonusProgramDAO model.ClientBonusProgram
	err := db.Where("program_id = ?", programID).First(&clientBonusProgramDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return clientBonusProgramDAO.ToDomain()
}

func (cbp *ClientBonusProgramRepository) GetAllWithClientID(
	conn abstract.IDBConnection, clientID int,
) ([]*domain.ClientBonusProgram, error) {
	db := conn.Get().(*gorm.DB)
	programs := []*model.ClientBonusProgram{}
	res := db.Where("client_id = ?", clientID).Find(programs)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	domainPrograms := []*domain.ClientBonusProgram{}
	for _, program := range programs {
		domainProg, err := program.ToDomain()
		if err != nil {
			return nil, err
		}

		domainPrograms = append(domainPrograms, domainProg)
	}

	return domainPrograms, nil
}
