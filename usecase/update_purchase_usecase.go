package usecase

import (
	"github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/purchase-api/gateway/txtdb"
)

//UpdatePurchaseUseCase save a domain.Purchase object
func UpdatePurchaseUseCase(ID string,it domain.Purchase) domain.Purchase {
	it = txtdb.UpdatePurchase(ID, it)
	UpdateItemPriceUseCase(it)
	return it
}

