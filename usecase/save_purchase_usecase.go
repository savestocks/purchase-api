package usecase

import (
	"github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/purchase-api/gateway/txtdb"
)

//SavePurchaseUseCase save a domain.Purchase object
func SavePurchaseUseCase(it domain.Purchase) domain.Purchase {
	it = txtdb.SavePurchase(it)
	UpdateItemPriceUseCase(it)
	return it
}

