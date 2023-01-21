package domain

import (
	"context"

	"prancerTest/internal/pancer/entity"
)

type PrancerUsecase interface {
	Calculate(ctx context.Context, transfer entity.Transfer) error
}
