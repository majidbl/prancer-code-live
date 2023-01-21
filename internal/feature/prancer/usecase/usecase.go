package usecase

import (
	"context"
	"errors"
	"fmt"
	"math"

	"prancerTest/internal/pancer/domain"
	"prancerTest/internal/pancer/entity"
)

type prancer struct {
	ID     int
	Name   string
	X      int
	Y      int
	Status entity.PrancerStatus
}

func New(id int) domain.PrancerUsecase {
	return &prancer{
		ID:     id,
		Name:   fmt.Sprintf("prancer-%d", id),
		X:      0,
		Y:      0,
		Status: entity.Running,
	}
}

func (p prancer) Calculate(ctx context.Context, transfer entity.Transfer) error {
	agentShortedPathName := ""
	shortedPath := 0

	for i := 0; i < entity.AgentCount; i++ {
		a, ok := entity.Agents.Load(fmt.Sprintf("agent-%d", i))
		if !ok {
			return errors.New("agent not found (invalid agent)")
		}

		agent := a.(domain.AgentUsecase)
		if agent.GetStatus(ctx) == entity.Moving {
			continue
		}

		distance := math.Sqrt(math.Pow(float64(transfer.XDestination-agent.GetX(ctx)), 2) + math.Pow(float64(transfer.YDestination-agent.GetY(ctx)), 2))
		if shortedPath == 0 {
			shortedPath = int(distance)
			agentShortedPathName = agent.GetName(ctx)
		} else if int(distance) < shortedPath {
			agentShortedPathName = agent.GetName(ctx)
		}
	}

	a, ok := entity.Agents.Load(agentShortedPathName)
	if !ok {
		return errors.New("agent not found (invalid agent)")
	}

	agent := a.(domain.AgentUsecase)
	agent.Move(ctx, entity.AgentTransfer{Transfer: transfer, Distance: shortedPath})

	return nil
}
