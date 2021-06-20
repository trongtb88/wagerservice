package wager

import (
	"context"
	"github.com/trongtb88/wagerservice/src/business/entity"
	"gorm.io/gorm"
	"time"
)

func (w *wager) createWagerSQL(ctx context.Context, param entity.WagerParam) (entity.Wager, error) {

	 wager := entity.Wager{
		TotalWagerValue:     param.TotalWagerValue,
		Odds:                param.Odds,
		SellingPercentage:   param.SellingPercentage,
		SellingPrice:        param.SellingPrice,
		CurrentSellingPrice: param.SellingPrice,
		PercentageSold:      0,
		AmountSold:          0,
		PlacedAt:            time.Now().UTC(),
	}

	tx := w.sql.Create(&wager)
	if tx.Error != nil {
		return wager, tx.Error
	}
	return wager, nil
}

func (w *wager) createBuyWagerSQL(ctx context.Context, param entity.BuyWagerParam) (entity.BuyWager, error) {

	buyWager := entity.BuyWager{
		WagerId:     param.WagerId,
		BuyingPrice: param.BuyingPrice,
		BoughtAt:    time.Now().UTC(),
	}

	wager, _ := w.getWagerByIdSQL(ctx, param.WagerId)
	wager.CurrentSellingPrice = param.BuyingPrice
	wager.AmountSold = wager.AmountSold + param.BuyingPrice
	wager.PercentageSold = (param.BuyingPrice/wager.CurrentSellingPrice) * 100

	err := w.sql.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&buyWager).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.Save(&wager).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return buyWager, err
	}
	return buyWager, nil
}

func (w *wager) getWagerByIdSQL(ctx context.Context, id int64) (entity.Wager, error) {

	var wager entity.Wager
	tx := w.sql.First(&wager, id)
	if tx.Error != nil {
		return wager, tx.Error
	}
	return wager, nil
}

func (w *wager) getWagersSQL(ctx context.Context, page int, limit int) ([]entity.Wager, error) {

	var wagers []entity.Wager

	offset := (page - 1) * limit
	tx := w.sql.Debug().Offset(offset).Limit(limit).Find(&wagers)

	if tx.Error != nil {
		return wagers, tx.Error
	}
	return wagers, nil
}




