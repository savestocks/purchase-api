package itemapi

import (

	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"os"
	"github.com/andersonlira/purchase-api/domain"
)

func UpdateItemPriceApi(ID string, itemPrice domain.ItemPrice, endpoint string) error {
	server := os.Getenv("PURCHASE_INTEGRATION_URL")
	user := os.Getenv("apikey")
	pass := os.Getenv("apisecret")
	
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s/price/%s", server,ID,endpoint)
	json, _ := json.Marshal(itemPrice)
	request, err := http.NewRequest("PATCH", url, strings.NewReader(string(json)))
	request.SetBasicAuth(user,pass)
	_, err = client.Do(request)
	if err != nil {
		return err
	} 
	return nil
}