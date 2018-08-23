package controller

import(
  "app/model"
  "net/http"
  "github.com/labstack/echo"
)

func GetSamples() echo.HandlerFunc {

  return func(c echo.Context) error {
    samples := []*model.Sample {
      model.NewSample(1,"hoge"),
      model.NewSample(2,"fuga"),
    }
    return c.JSON(http.StatusOK, samples)
  }
}

