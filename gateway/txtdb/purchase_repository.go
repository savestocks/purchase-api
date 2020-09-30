package txtdb

import (
	"encoding/json"
	"errors"
    "fmt"
	"log"
	"time"

    "github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/goutils/str"
)

//GetPurchaseList return all items 
func GetPurchaseList(itemID string) []domain.Purchase {
	list := []domain.Purchase{}
    fileName := fmt.Sprintf("bd/%s.json", itemID);
	listTxt, _ := io.ReadFile(fileName)
	json.Unmarshal([]byte(listTxt), &list)
	return list
}

//GetPurchaseByID return all items 
func GetPurchaseByID(ID string,itemID string) (domain.Purchase, error) {
	list := GetPurchaseList(itemID)
	for idx, _ := range list {
		if(list[idx].ID == ID){
			return list[idx],nil
		}
	}
	return domain.Purchase{}, errors.New("NOT_FOUND")
}



//SavePurchase saves a Purchase object
func SavePurchase(it domain.Purchase) domain.Purchase {
	list := GetPurchaseList(it.ItemID)
	it.ID = str.NewUUID()
	it.CreatedAt = time.Now()
	list = append(list, it)
	writePurchase(list)
	return it
}

//UpdatePurchase( updates a Purchase object
func UpdatePurchase(ID string, it domain.Purchase) domain.Purchase{
	list := GetPurchaseList(it.ItemID)
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list[idx] = it
			list[idx].ID = ID
			list[idx].UpdatedAt = time.Now()
			writePurchase(list)
			return list[idx]
		}
	}
	return it
}

//DeletePurchase delete object by giving ID
func DeletePurchase(ID string,itemID string) bool {
	list := GetPurchaseList(itemID)
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list = append(list[:idx], list[idx+1:]...)
			writePurchase(list)
			return true
		}
	}
	return false
}

func writePurchase(list []domain.Purchase) {
	if len(list) == 0 {
		return 
	}
	itemID := list[0].ItemID
	b, err := json.Marshal(list)
	if err != nil {
		log.Println("Error while writiong file items")
		return
	}
	io.WriteFile(fmt.Sprintf("bd/%s.json", itemID), string(b))
}
