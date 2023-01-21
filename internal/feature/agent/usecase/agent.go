package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"prancerTest/internal/pancer/domain"
	"prancerTest/internal/pancer/entity"
)

type agent struct {
	ID     int
	Name   string
	X      int
	Y      int
	Status entity.Status
}

func New(id int) domain.AgentUsecase {
	return &agent{
		ID:     id,
		Name:   fmt.Sprintf("agent-%d", id),
		X:      0,
		Y:      0,
		Status: entity.Stop,
	}
}

func (a *agent) Move(ctx context.Context, transfer entity.AgentTransfer) {
	a.Status = entity.Moving
	for i := transfer.Distance; i >= 0; i-- {
		log.Printf("%s distance form prancer : %d", a.GetName(ctx), i)
		time.Sleep(time.Second * 1)
	}

	if transfer.XDestination > a.X {
		a.X += transfer.XDestination
	} else if transfer.XDestination < a.X {
		a.X -= transfer.XDestination
	}

	if transfer.YDestination > a.Y {
		a.Y += transfer.YDestination
	} else if transfer.YDestination < a.Y {
		a.Y -= transfer.YDestination
	}

	a.Status = entity.Stop
}

func (a *agent) GetName(ctx context.Context) string {
	if a == nil {
		return ""
	}

	return a.Name
}

func (a *agent) GetX(ctx context.Context) int {
	if a == nil {
		return 0
	}

	return a.X
}

func (a *agent) GetY(ctx context.Context) int {
	if a == nil {
		return 0
	}

	return a.Y
}

func (a *agent) GetStatus(ctx context.Context) entity.Status {
	if a == nil {
		return ""
	}

	return a.Status
}
