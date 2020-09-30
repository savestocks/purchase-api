package controller

import (
    "net/http"

    "github.com/andersonlira/purchase-api/gateway/txtdb"
    "github.com/andersonlira/purchase-api/domain"
	"github.com/labstack/echo/v4"

)


//GetPurchaseList return all objects 
func GetPurchaseList(c echo.Context) error {
    itemID := c.Param("itemId")

    list := txtdb.GetPurchaseList(itemID)

	return c.JSON(http.StatusOK, list)
}

func GetPurchaseByID(c echo.Context) error {
    ID := c.Param("id")
    itemID := c.Param("itemId")
    it, err := txtdb.GetPurchaseByID(ID,itemID)
    if err != nil {
        return c.JSON(http.StatusNotFound,it)
    }
    return c.JSON(http.StatusOK, it)
}

func SavePurchase(c echo.Context) error {
    it := domain.Purchase{}
    c.Bind(&it)
    it = txtdb.SavePurchase(it)
    return c.JSON(http.StatusCreated, it)
}

func UpdatePurchase(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Purchase{}
    c.Bind(&it)
    it = txtdb.UpdatePurchase(ID,it)
    return c.JSON(http.StatusOK, it)
}

func DeletePurchase(c echo.Context) error {
    ID := c.Param("id")
    itemID := c.Param("itemId")

    result := txtdb.DeletePurchase(ID,itemID)
    return c.JSON(http.StatusOK, result)
}