package route

import (
  "app/api/controller/v1"
  "github.com/labstack/echo"
)

func Init() *echo.Echo {
	
  e := echo.New()

  api := e.Group("/api/v1")
  {
    api.GET("/samples", controller.GetSamples())
    api.GET("/accounts", controller.GetAccounts())
    api.POST("/accounts", controller.PostAccounts())
    api.PUT("/accounts", controller.PutAccounts())
    api.DELETE("/accounts", controller.DeleteAccounts())
    api.GET("/heartBeat", controller.GetHeartBeat())
  }
  return e
}
