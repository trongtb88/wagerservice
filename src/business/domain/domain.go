package domain

import (
	"github.com/trongtb88/wagerservice/src/business/domain/wager"
	"gorm.io/gorm"
)

type Domain struct {
	Wager       wager.DomainItf
}

func Init(
	sql *gorm.DB,
	) *Domain {

	return &Domain{
		Wager: wager.InitWagerDomain(
			sql,
		),
	}
}
