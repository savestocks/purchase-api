package usecase

import (
	"fmt"
	//"github.com/andersonlira/purchase-api/domain"
	//"github.com/andersonlira/purchase-api/gateway/txtdb"
)

//SavePurchaseUseCase save a domain.Purchase object
func RemoveOldPurchaseUseCase(IDS []string) bool {

	for _, ID := range IDS {
		fmt.Println(ID)
	}

	//it = txtdb.GetPurchaseList(it)
	//UpdateItemPriceUseCase(it)
	return true
}

