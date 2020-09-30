package controller

import (
	"os"

	"github.com/andersonlira/purchase-api/config"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//MapRoutes for the endpoints which the API listens for
func MapRoutes(e *echo.Echo) {
	g := e.Group("/purchase-api/v1")
	if config.Values.UsePrometheus {
		p := prometheus.NewPrometheus("echo", nil)
		p.Use(e)
	}
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		key := os.Getenv("apikey")
		secret := os.Getenv("apisecret")
		if username == key && password == secret {
			return true, nil
		}
		return false, nil
	}))	

	g.GET("/purchase/:itemId", GetPurchaseList)
	g.GET("/purchase/:itemId/:id", GetPurchaseByID)
	g.POST("/purchase", SavePurchase)
	g.PUT("/purchase/:id", UpdatePurchase)
	g.DELETE("/purchase/:itemId/:id", DeletePurchase)
	g.GET("/health", CheckHealth)
	g.GET("/info", GetInfo)
}
