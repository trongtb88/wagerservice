package wager

import (
	"context"
	Wager "github.com/trongtb88/wagerservice/src/business/domain/wager"
	"github.com/trongtb88/wagerservice/src/business/entity"
)

// Usecaseitf uc insterface
type Usecaseitf interface {
	CreateWager(ctx context.Context, param entity.WagerParam) (entity.Wager, error)
	GetWagerById(ctx context.Context, id int64) (entity.Wager, error)
	CreateBuyWager(ctx context.Context, param entity.BuyWagerParam) (entity.BuyWager, error)
	GetWagers(ctx context.Context, page int, limit int) ([]entity.Wager, error)
}

type wager struct {
	wager Wager.DomainItf
}

// Options for uc
type Options struct {
}

// InitLogistic init logistic uc
func InitWager(
	Wager Wager.DomainItf,
) Usecaseitf {
	return &wager{
		Wager,
	}
}
