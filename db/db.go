package db

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/array?parseTime=true")

  if err != nil {
    fmt.Println("db is not connected")
		panic(err.Error())
	} else {
		fmt.Println("db is connected")
	}

  // err = db.Ping()

  // fmt.Println(err)
	// if err != nil {
	// 	fmt.Println("db is not connected")
	// 	panic(err.Error())
	// }
  return db

  // defer db.Close()
}
