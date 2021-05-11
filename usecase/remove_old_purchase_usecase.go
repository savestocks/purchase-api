package usecase

import (
	"regexp"
	"time"
	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/purchase-api/gateway/txtdb"
)

//SavePurchaseUseCase save a domain.Purchase object
func RemoveOldPurchaseUseCase() bool {
	sixMonthsAgo := time.Now().AddDate(-1,0,0)

	for _, ID := range getIDS() {
		txtdb.DeleteOld(ID,sixMonthsAgo)
	}

	return true
}

func getIDS() (IDS []string) {
	files, _ := io.ListFiles("bd/",[]string{"json"})
	for _, file := range files {
		re := regexp.MustCompile(`(.*)[/ | \\](.*)(\.json)`)
		ID := re.ReplaceAllString(file, "$2")
		IDS = append(IDS, ID)
	}
	return
}