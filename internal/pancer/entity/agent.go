package entity

type Status string

const (
	Started Status = "started"
	Stop    Status = "stop"
	Moving  Status = "moving"
)

type Agent struct {
	ID     int
	Name   string
	X      int
	Y      int
	Status Status
}

type AgentTransfer struct {
	Transfer
	Distance int
}
