package main

import (
  "app/route"
)

func main() {
  e := route.Init()
  e.Logger.Fatal(e.Start(":5000"))
}
