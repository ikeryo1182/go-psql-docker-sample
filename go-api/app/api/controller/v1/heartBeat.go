package controller

import(
  "app/model"
  "net/http"
  "github.com/labstack/echo"
)

func GetHeartBeat() echo.HandlerFunc {
  return func(c echo.Context) error {
    heartBeat := model.NewHeartBeat("1.0.0", "Good Heart")
    return c.JSON(http.StatusOK, heartBeat)
  }
}