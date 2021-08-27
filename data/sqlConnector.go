package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

//Structure to pass the SQL connection parameters
type SqlParams struct {
	Server   string
	Port     int
	User     string
	Passwd   string
	Database string
}

//sql context to call the queries. Should have been an external configuration
var sqlParams = SqlParams{
	Server:   "DFDEVNBMD33",
	Port:     1433,
	User:     "go_sql",
	Passwd:   "go_sql123",
	Database: "WalletDA",
}

//encapsulates connection to a DB
func SqlConnect() (*sql.DB, error) {

	var db *sql.DB
	//connString := fmt.Sprintf("sqlserver://go_sql:go_sql123@DFDEVNBMD33?database=WalletDA")
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		sqlParams.User, sqlParams.Passwd, sqlParams.Server, sqlParams.Port, sqlParams.Database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
		return db, err
	}
	log.Printf("Connected to SQL database!\n")
	return db, nil
}

//database interactions for the withdraw and deposit endpoints
func ExecuteTransaction(currency, user string, amount float64) error {
	//try to connect
	db, err := SqlConnect()
	if err != nil {
		//move the error forward
		return err
	}
	defer db.Close()
	//execute the stored procedure
	//empty context
	ctx := context.TODO()
	var result string
	_, errSp := db.ExecContext(ctx, "spAddTransaction",
		sql.Named("UserName", user),
		sql.Named("Currency", currency),
		sql.Named("Amount", amount),
		sql.Named("returnMessage", sql.Out{Dest: &result}),
	)
	if errSp != nil {
		fmt.Println("Error executing transaction: " + err.Error())
		return errSp
	}
	//let the caller act on the message
	return errors.New(result)
}

//database interactions for the history endpoint
func GetHistorySQL(user string, startDate, endDate time.Time) (Transactions, error) {
	//try to connect
	db, err := SqlConnect()
	var HistoryResult Transactions
	if err != nil {
		//move the error forward
		return HistoryResult, err
	}
	defer db.Close()
	//run the query
	//empty context
	ctx := context.TODO()
	rows, errQuery := db.QueryContext(ctx, `SELECT Currency, Amount, TimeStamp FROM TblTransactions join TblUsers on TblUsers.UserId = TblTransactions.userId where UserName = @p1 and TimeStamp between @p2 and @p3;`, user, startDate, endDate)
	if errQuery != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return HistoryResult, err
	}
	defer rows.Close()

	for rows.Next() {
		var element TransactionHistory
		errScan := rows.Scan(&element.Currency, &element.Amount, &element.TimeStamp)
		if errScan != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return HistoryResult, err
		}
		//append to historyresult
		HistoryResult = append(HistoryResult, &element)

	}
	return HistoryResult, nil
}

func GetUserBalanceSQL(userName string) (UserBalance, error) {

	//Get USerBalance from database
	//try to connect
	db, err := SqlConnect()
	var userBalance UserBalance
	userBalance.UserName = userName
	if err != nil {
		//move the error forward
		return userBalance, err
	}
	defer db.Close()
	//run the query
	//empty context
	ctx := context.TODO()
	rows, errQuery := db.QueryContext(ctx, `SELECT Currency,Amount FROM TblBalance join TblUsers on TblUsers.UserId = TblBalance.UserId where UserName = @p1;`, userName)
	if errQuery != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return userBalance, err
	}
	defer rows.Close()

	for rows.Next() {
		var element CryptoBalance
		errScan := rows.Scan(&element.Currency, &element.Amount)
		if errScan != nil {
			fmt.Println("Error reading rows: " + errScan.Error())
			return userBalance, errScan
		}
		//append to historyresult
		userBalance.CryptoBalanceList = append(userBalance.CryptoBalanceList, element)
	}

	//calculate user balance and return
	result, errCalc := userBalance.CalculateUserBalance()
	if errCalc != nil {
		fmt.Println("Error reading rows: " + errCalc.Error())
		return userBalance, errCalc
	}
	return result, nil

}
