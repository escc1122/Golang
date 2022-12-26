package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//https://ithelp.ithome.com.tw/articles/10243865

func getConn() *sql.DB {
	db, err := sql.Open("mysql", "test:12345@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func PrintStats(msg string, db *sql.DB) {
	stats := db.Stats()

	fmt.Printf("%s %s : %v \n", msg, "MaxOpenConnections", stats.MaxOpenConnections)
	fmt.Printf("%s %s : %v \n", msg, "InUse", stats.InUse)
	fmt.Printf("%s %s : %v \n", msg, "OpenConnections", stats.OpenConnections)
	fmt.Printf("%s %s : %v \n", msg, "WaitCount", stats.WaitCount)
	fmt.Println("")
}

/*
Ping() MaxOpenConnections : 10
Ping() InUse : 0
Ping() OpenConnections : 1
Ping() WaitCount : 0
*/
func TestPing() {
	db := getConn()
	PrintStats("getConn()", db)

	//db.Ping() 呼叫完畢後會馬上把連線返回給連線池
	db.Ping()
	PrintStats("Ping()", db)
}

/*
stmt.Exec MaxOpenConnections : 10
stmt.Exec InUse : 0
stmt.Exec OpenConnections : 1
stmt.Exec WaitCount : 0
*/
func TestPrepare() {
	db := getConn()

	//db.Exec() 呼叫完畢後會馬上把連線返回給連線池，但是它返回的Result物件還保留這連線的引用，當後面的程式碼需要處理結果集的時候連線將會被重用。
	stmt, _ := db.Prepare("SELECT 1")
	PrintStats("db.Prepare", db)

	stmt.Exec("")
	PrintStats("stmt.Exec", db)
}

/*
db.Query MaxOpenConnections : 10
db.Query InUse : 1
db.Query OpenConnections : 1
db.Query WaitCount : 0

row.Close() MaxOpenConnections : 10
row.Close() InUse : 0
row.Close() OpenConnections : 1
row.Close() WaitCount : 0
*/
func TestQuery() {
	db := getConn()

	//db.Query() 呼叫完畢後會將連線傳遞給sql.Rows型別，當然後者迭代完畢或者顯示的呼叫.Close()方法後，連線將會被釋放回到連線池(<---踩雷點，如果沒close掉是不會釋放掉連線)
	row, _ := db.Query("SELECT 1")

	PrintStats("db.Query", db)
	row.Close()
	PrintStats("row.Close()", db)
}

/*
db.QueryRow MaxOpenConnections : 10
db.QueryRow InUse : 1
db.QueryRow OpenConnections : 1
db.QueryRow WaitCount : 0

row2.Scan MaxOpenConnections : 10
row2.Scan InUse : 0
row2.Scan OpenConnections : 1
row2.Scan WaitCount : 0
*/
func TestQueryRow() {
	db := getConn()

	//db.QueryRow()呼叫完畢後會將連線傳遞給sql.Row型別，當.Scan()方法呼叫之後把連線釋放回到連線池。
	row2 := db.QueryRow("SELECT 1 as a")
	PrintStats("db.QueryRow", db)
	var a interface{}

	switch err := row2.Scan(&a); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("")
	default:
		panic(err)
	}
	PrintStats("row2.Scan", db)
}

/*
db.Begin() before MaxOpenConnections : 10
db.Begin() before InUse : 0
db.Begin() before OpenConnections : 0
db.Begin() before WaitCount : 0

db.Begin() MaxOpenConnections : 10
db.Begin() InUse : 1
db.Begin() OpenConnections : 1
db.Begin() WaitCount : 0

tx.Prepare MaxOpenConnections : 10
tx.Prepare InUse : 1
tx.Prepare OpenConnections : 1
tx.Prepare WaitCount : 0

tx stmt.Exec MaxOpenConnections : 10
tx stmt.Exec InUse : 1
tx stmt.Exec OpenConnections : 1
tx stmt.Exec WaitCount : 0

tx.Commit() MaxOpenConnections : 10
tx.Commit() InUse : 0
tx.Commit() OpenConnections : 1
tx.Commit() WaitCount : 0

stmt.Close() MaxOpenConnections : 10
stmt.Close() InUse : 0
stmt.Close() OpenConnections : 1
stmt.Close() WaitCount : 0

tx.Rollback() MaxOpenConnections : 10
tx.Rollback() InUse : 0
tx.Rollback() OpenConnections : 1
tx.Rollback() WaitCount : 0
*/
func TestBegin() {
	db := getConn()

	PrintStats("db.Begin() before", db)

	tx, err := db.Begin()
	PrintStats("db.Begin()", db)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("SELECT 1 as a")
	PrintStats("tx.Prepare", db)
	if err != nil {
		log.Fatal(err)
	}

	// test1
	_, err = stmt.Exec("")
	PrintStats("tx stmt.Exec", db)

	//test2
	err = tx.Commit()
	PrintStats("tx.Commit()", db)

	//if err != nil {
	//	log.Fatal(err)
	//}
	//test3
	stmt.Close()
	PrintStats("stmt.Close()", db)

	//test4
	tx.Rollback()
	PrintStats("tx.Rollback()", db)
}

func main() {
	TestPing()
	TestPrepare()
	TestQuery()
	TestQueryRow()
	TestBegin()
}
