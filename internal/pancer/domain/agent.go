package domain

import (
	"context"

	"prancerTest/internal/pancer/entity"
)

type AgentUsecase interface {
	Move(ctx context.Context, transfer entity.AgentTransfer)
	GetName(ctx context.Context) string
	GetStatus(ctx context.Context) entity.Status
	GetX(ctx context.Context) int
	GetY(ctx context.Context) int
}
