package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

//SQL tests
type SqlParams struct {
	Server string
	Port   int
	User   string
	Passwd string
}

func SqlTest(sqlParams SqlParams) {
	// var server = "DFDEVNBMD33"
	// var port = 1433
	var db *sql.DB
	//connString := fmt.Sprintf("sqlserver://go_sql:go_sql123@DFDEVNBMD33?database=WalletDA")
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=WalletDA", sqlParams.User, sqlParams.Passwd, sqlParams.Server, sqlParams.Port)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")
	defer db.Close()

	tsql := fmt.Sprintf("SELECT UserId,Currency,Amount,TimeStamp FROM dbo.TblTransactions")
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var UserId, Currency, TimeStamp string
		var Amount float64
		err := rows.Scan(&UserId, &Currency, &Amount, &TimeStamp)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return
		}
		fmt.Println("User: %d, Currency: %s, Amount: %d TimeStamp %s\n", UserId, Currency, Amount, TimeStamp)
		count++
	}
	fmt.Println(count)
}
