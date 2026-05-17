package abstract

type IClientJoinService interface {
	JoinProgram(clientID int, programID int) (*int, error)
}
