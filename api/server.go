package api

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/SaidovZohid/market-project/api/v1"
	"github.com/SaidovZohid/market-project/config"
	"github.com/SaidovZohid/market-project/storage"
)

type RouteOptions struct {
	Cfg *config.Config
	Storage storage.StorageI
}

func New(opt *RouteOptions) *gin.Engine {
	router := gin.Default()

	handler := v1.New(&v1.HandlerV1{
		Cfg: opt.Cfg,
		Storage: &opt.Storage,
	})

	apiV1 := router.Group("/debts")
	{
		apiV1.POST("", handler.CreateDebt)
		apiV1.GET("/:id", handler.GetDebt)
		apiV1.PUT("/:id", handler.UpdateDebt)
		apiV1.DELETE("/:id", handler.DeleteDebt)
		apiV1.GET("/all", handler.GetAllDebts)
	}
	return router
}