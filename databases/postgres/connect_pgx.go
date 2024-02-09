package main

import (
  "database/sql"
  "net/url"

  _ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
  dsn := url.URL{
    Scheme: "postgres",
    Host: "localhost:5432",
    User: url.UserPassword("user", "password"),
    Path: "dbname",
  }

  q := dsn.Query()
  q.Add("sslmode", "disable")

  dsn.RawQuery = q.Encode()

  db, err := sql.Open("pgx", dsn.String())
  if err != nil {
    fmt.Println("sql.Open", err)
    return
  }
  defer func(){
    _ = db.Close()
    fmt.Println("closed")
  }()
}
