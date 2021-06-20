package entity

import "time"

type Wager struct {
	Id                  int64     `gorm:"primaryKey;autoIncrement;not_null" json:"id"`
	TotalWagerValue     int64     `gorm:"type:bigint(20);" json:"total_wager_value"`
	Odds                int64     `gorm:"type:bigint(20);" json:"odds"`
	SellingPercentage   int       `gorm:"type:int(3);" json:"selling_percentage"`
	SellingPrice        float64   `gorm:"type:decimal(10,3);" json:"selling_price"`
	CurrentSellingPrice float64   `gorm:"type:decimal(10,3);" json:"current_selling_price"`
	PercentageSold      float64   `gorm:"type:int(3);" json:"percentage_sold"`
	AmountSold          float64   `gorm:"type:decimal(10,3);" json:"amount_sold"`
	PlacedAt           time.Time  `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"placed_at"`
}

type BuyWager struct {
	Id                  int64     `gorm:"primaryKey;autoIncrement;not_null" json:"id"`
	WagerId             int64     `gorm:"type:bigint(20);" json:"total_wager_value"`
	BuyingPrice        float64    `gorm:"type:decimal(10,3);" json:"selling_price"`
	BoughtAt           time.Time  `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"bought_at"`
}

type WagerParam struct {
	TotalWagerValue int64  `json:"total_wager_value"`
	Odds int64             `json:"odds"`
	SellingPercentage int  `json:"selling_percentage"`
	SellingPrice float64   `json:"selling_price"`
}

type BuyWagerParam struct {
	WagerId     int64
	BuyingPrice float64 `json:"buying_price"`
}
