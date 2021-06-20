package wager

import (
	"context"
	"github.com/trongtb88/wagerservice/src/business/entity"
)

func (w *wager) CreateWager(ctx context.Context, param entity.WagerParam) (entity.Wager, error) {
	return w.wager.CreateWager(ctx, param)
}

func (w *wager) GetWagerById(ctx context.Context, id int64) (entity.Wager, error) {
	return w.wager.GetWagerById(ctx, id)
}

func (w *wager) CreateBuyWager(ctx context.Context, param entity.BuyWagerParam) (entity.BuyWager, error) {
	return w.wager.CreateBuyWager(ctx, param)
}

func (w *wager) GetWagers(ctx context.Context, page int, limit int) ([]entity.Wager, error) {
	return w.wager.GetWagers(ctx, page, limit)
}
