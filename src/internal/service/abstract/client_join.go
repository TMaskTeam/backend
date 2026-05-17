package abstract

import (
	"backend/src/internal/domain"
)

type IClientJoinService interface {
	Join(client *domain.Client, programID int) (int, int, error)
}
