package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"io/ioutil"

	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/purchase-api/domain"
	"github.com/andersonlira/purchase-api/usecase"
)

var source string

var markets = map[string]string{
	"Continente":"4a26b45b-7fb2-1fe1-ad87-629d6c83b081",
	"Decathlon":"41ed62f1-80b9-cc79-931d-b5bb755842eb",
	"E'leclerc":"2e4f651d-ee14-2a9b-6385-2fab9b6c9e65",
	"Intermarché":"73528e21-2f80-880b-d9df-a42238e7f6e4",
	"Lidl":"bec302b0-0640-f486-67ee-86166e523647",
	"Minipreço":"e3266429-5a30-3181-758d-0e4d22c1e21f",
	"Outro":"ed34a30c-804e-7a00-6a4b-23163b64a534",
	"Pingo Doce":"6cb9cfcb-67cd-7a8b-7c7c-797bd573f9bb",
}

var wg sync.WaitGroup

func main(){

	fmt.Println("Starting migration Purcahse")
	if len(os.Args) != 2{
		panic("usage: cmd <source>")
	}

	source = os.Args[1]

	files := getFiles()

	for idx,file := range files {
		fmt.Printf("Processing %d of %d - %s\n",idx+1,len(files),file)
		purchases := getPurchases(file)

		for _, p := range purchases {
			wg.Add(1)
			purchase := domain.Purchase{}
			purchase.CreatedAt = p.When
			purchase.ItemID = file
			purchase.Price = p.Price
			
			marketID, _ := markets[p.Market]

			if marketID == "" {
				marketID = markets["Outro"]
			}
			purchase.MarketID = marketID
			purchase.Qtd = float32(p.Qtd)
			go worker(purchase)
		}


	}
	wg.Wait()

}

func worker(purchase domain.Purchase){
	defer wg.Done()
	usecase.SavePurchaseUseCase(purchase)
}

type purchase struct {
	ID     string    `json:"id"`
	When   time.Time `json:"when"`
	Qtd    int32     `json:"qtd"`
	Price  int32     `json:"price"`
	Market string    `json:"market"`
}

func getPurchases(ID string) []purchase {
	purchases := []purchase{}
	listTxt, _ := io.ReadFile(fmt.Sprintf("%s/%s.json",source, ID))
	json.Unmarshal([]byte(listTxt), &purchases)
	return purchases
}


func getFiles() (list []string) {
    files, err := ioutil.ReadDir(source)
    if err != nil {
    	panic(err)
    }
	
    for _, f := range files {
			list = append(list,strings.ReplaceAll(f.Name(),".json",""))
	}
	return list	
}