package usecase 

import (
	"github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/purchase-api/gateway/itemapi"
	"github.com/andersonlira/purchase-api/gateway/txtdb"
)

func UpdateItemPriceUseCase(purchase domain.Purchase) error {
	list := txtdb.GetPurchaseList(purchase.ItemID)	
	low, high := getStatistics(list,purchase.Price)
	var err error
	if  purchase.Price <= low {
		err =  itemapi.UpdateItemPriceApi(domain.ItemPrice{purchase.ID,purchase.Price},"lowest")
	}
	if err == nil && purchase.Price >= high{
		err = itemapi.UpdateItemPriceApi(domain.ItemPrice{purchase.ID,purchase.Price},"highest")
	}
	if err == nil {
		err = itemapi.UpdateItemPriceApi(domain.ItemPrice{purchase.ID,purchase.Price},"last")
	}

	return err
}

 

func getStatistics(items []domain.Purchase, last int32) (low int32, high int32) {
	for _, it := range items {
		if low == 0 || low > it.Price {
			low = it.Price
		}
		if high == 0 || high < it.Price {
			high = it.Price
		}
	}
	if last < low {
		low = last
	}
	if last > high {
		high = last
	}

	if low == 0 {
		low = last
	}

	return
}
