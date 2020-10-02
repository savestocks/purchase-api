package usecase 

import (
	"errors"
	"github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/purchase-api/gateway/itemapi"
)

func UpdateItemPrice(itemPrice domain.ItemPrice) error {
	return itemapi.UpdateItemPriceApi(itemPrice)
}
