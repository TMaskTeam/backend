package public

import (
	context "backend/src/internal/context/abstract"
	"backend/src/internal/dto"
	"backend/src/internal/service/abstract"
	"errors"
	"net/http"
	"strconv"
)

func ClientJoinProgram(
	ctx context.HandlerContext,
	request *dto.ClientJoinProgramRequest,
	clientJoinService abstract.IClientJoinService,
	clientService abstract.IClientService,
) (interface{}, error) {
	userID, ok := ctx.GetLocal("user_id").(int)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	role, ok := ctx.GetLocal("role").(string)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	if role == "client" {
		client, err := clientService.GetByID(userID)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return nil, err
		}
		programID := ctx.Params("program_id")
		id, err := strconv.Atoi(programID)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return dto.ClientJoinProgramResponse{}, err
		}

		clientBonusProgramID, err := clientJoinService.JoinProgram(client.ID, id)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return dto.ClientJoinProgramResponse{}, err
		}

		resp := buildClientJoinProgramResponse(*clientBonusProgramID)

		ctx.Status(http.StatusOK)
		return resp, nil

	} else {
		ctx.Status(http.StatusBadRequest)
		return nil, errors.New("unknown role")
	}
}

func buildClientJoinProgramResponse(clientBonusProgramID int) dto.ClientJoinProgramResponse {
	return dto.ClientJoinProgramResponse{
		ClientBonusProgramID: clientBonusProgramID,
	}
}
