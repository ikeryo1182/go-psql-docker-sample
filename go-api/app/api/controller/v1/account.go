package controller

import(
  "log"
  "strconv"
  "app/model"
  "net/http"
  "github.com/labstack/echo"
    
  "database/sql"
  _ "github.com/lib/pq"
)

func GetAccounts() echo.HandlerFunc {
  return func(c echo.Context) error {
    db, err := sql.Open("postgres", "postgresql://root:admin1234@go-api-db/goapi?sslmode=disable")
    defer db.Close()
    if err != nil {
      // エラー時の処理
      log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM account;")
    if err != nil {
      // エラー時の処理
      log.Fatal(err)
    }
  
    var accounts []model.Account
    for rows.Next() {
      var account model.Account
      rows.Scan(&account.Id, &account.Name)
      accounts = append(accounts, account)
    }
  
    return c.JSON(http.StatusOK, accounts)
  }
}

func PostAccounts() echo.HandlerFunc {
  return func(c echo.Context) error {
    name := c.QueryParam("name")
    db, err := sql.Open("postgres", "postgresql://root:admin1234@go-api-db/goapi?sslmode=disable")
    defer db.Close()
    if err != nil {
      // エラー時の処理
      log.Fatal(err)
    }

    _, err = db.Exec("INSERT INTO account VALUES (nextval('id_seq'), $1);", name)
    if err != nil {
      // エラー時の処理
      log.Fatal(err)
    }
    return c.JSON(http.StatusOK, err)
  }
}

func PutAccounts() echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _  := strconv.Atoi(c.QueryParam("id"))
    name := c.QueryParam("name")
    db, err := sql.Open("postgres", "postgresql://root:admin1234@go-api-db/goapi?sslmode=disable")
    defer db.Close()
    if err != nil {
      // エラー時の処理
      log.Fatal(err)
    }
    
    _, err = db.Exec("UPDATE account SET name = $1 WHERE id = $2", name, id)
    if err != nil {
      // エラー時の処理
      log.Fatal(err)
    }

    return c.JSON(http.StatusOK, id)
  }
}

func DeleteAccounts() echo.HandlerFunc {

  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.QueryParam("id"))
    db, err := sql.Open("postgres", "postgresql://root:admin1234@go-api-db/goapi?sslmode=disable")
    defer db.Close()
    if err != nil {
      // エラー時の処理
      log.Fatal(err)
    }
    
    _, err = db.Exec("DELETE FROM account WHERE id = $1;", id)
    if err != nil {
      // エラー時の処理
      log.Fatal(err)
    }
    return c.JSON(http.StatusOK, id)
  }
}



