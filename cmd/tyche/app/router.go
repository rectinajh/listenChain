package app

import (
	"ethgo/util/ginx"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func (app *App) Router(g *gin.Engine) {
	g.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	g.POST("/tyche/api/transact", ginx.WrapHandler(app.Transact))
	g.POST("/tyche/api/call", ginx.WrapHandler(app.Call))
	g.POST("/tyche/api/wallet/create", ginx.WrapHandler(app.Create))
	g.POST("/tyche/api/wallet/balance_at", ginx.WrapHandler(app.BalanceAt))
	g.POST("/tyche/api/wallet/minter", ginx.WrapHandler(app.Minter))
	g.POST("/tyche/api/wallet/sign", ginx.WrapHandler(app.Sign))
	g.POST("/tyche/api/wallet/UserSignHash", ginx.WrapHandler(app.UserSignHash))
	g.POST("/tyche/api/order/get", ginx.WrapHandler(app.Order))
	g.POST("/tyche/api/wallet/contractTxCount", ginx.WrapHandler(app.ContractTxCount))
	g.POST("/tyche/api/wallet/GetGasPrice", ginx.WrapHandler(app.GetGasPrice))
	g.POST("/tyche/api/wallet/IsContractAddress", ginx.WrapHandler(app.IsContractAddress))
	g.POST("/tyche/api/wallet/CompareBytecodeAndSourceCode", ginx.WrapHandler(app.CompareBytecodeAndSourceCode))
}
