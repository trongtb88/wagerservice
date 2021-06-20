package wager

import (
	"context"
	"github.com/trongtb88/wagerservice/src/business/entity"
	"gorm.io/gorm"
)

// DomainItf domain interface for logistic
type DomainItf interface {
	CreateWager(ctx context.Context, param entity.WagerParam) (entity.Wager, error)
	GetWagerById(ctx context.Context, id int64) (entity.Wager, error)
	CreateBuyWager(ctx context.Context, param entity.BuyWagerParam) (entity.BuyWager, error)
	GetWagers(ctx context.Context, page int, limit int) ([]entity.Wager, error)
}

type wager struct {
	sql         *gorm.DB
}

// InitLogisticDomain logistic domain init
func InitWagerDomain(
	sql *gorm.DB,
	) DomainItf {
	a := &wager{
		sql:         sql,
	}
	return a
}
