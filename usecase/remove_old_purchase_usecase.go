package usecase

import (
	"time"
	//"github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/purchase-api/gateway/txtdb"
)

//SavePurchaseUseCase save a domain.Purchase object
func RemoveOldPurchaseUseCase(IDS []string) bool {

	sixMonthsAgo := time.Now().AddDate(0,-6,0)

	for _, ID := range IDS {
		txtdb.DeleteOld(ID,sixMonthsAgo)
	}

	//it = txtdb.GetPurchaseList(it)
	//UpdateItemPriceUseCase(it)
	return true
}

