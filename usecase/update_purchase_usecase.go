package usecase

import (
	"github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/purchase-api/gateway/txtdb"
)

//UpdatePurchaseUseCase save a domain.Purchase object
func UpdatePurchaseUseCase(ID string,it domain.Purchase) domain.Purchase {
	return txtdb.UpdatePurchase(ID, it)
}

