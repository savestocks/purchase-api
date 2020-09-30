package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/andersonlira/goutils/io"
)

var source string
var target string

func main(){
	
	fmt.Println("Starting migration")
	if len(os.Args) != 3{
		panic("usage: cmd <source> <target>")
	}

	source = os.Args[1]
	target = os.Args[2]

	purchases := getPurchases("f8ea6fdf-3846-bfca-a739-b8b3ed62b15d")

	for _, p := range purchases {
		fmt.Println(p.ID, p.Market)
	}
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