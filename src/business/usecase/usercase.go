package usecase

import (
	"github.com/trongtb88/wagerservice/src/business/domain"
	"github.com/trongtb88/wagerservice/src/business/usecase/wager"
)

type Usecase struct {
	Wager wager.Usecaseitf
}

// Init init all usecase
func Init(
	dom *domain.Domain,
) *Usecase {

	return &Usecase{
		Wager: wager.InitWager(
			dom.Wager,
		),
	}
}
