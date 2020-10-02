package usecase

import (
	"github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/purchase-api/gateway/txtdb"
)

//SavePurchaseUseCase save a domain.Purchase object
func SavePurchaseUseCase(it domain.Purchase) domain.Purchase {
	return txtdb.SavePurchase(it)
}

