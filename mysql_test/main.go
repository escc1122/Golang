package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
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

func main() {

	db := getConn()
	PrintStats("getConn()", db)

	//db.Ping() 呼叫完畢後會馬上把連線返回給連線池
	db.Ping()
	PrintStats("Ping()", db)

	//db.Exec() 呼叫完畢後會馬上把連線返回給連線池，但是它返回的Result物件還保留這連線的引用，當後面的程式碼需要處理結果集的時候連線將會被重用。
	stmt, _ := db.Prepare("SELECT 1")
	PrintStats("db.Prepare", db)

	res, _ := stmt.Exec("")
	PrintStats("stmt.Exec", db)
	fmt.Println(res)

	//db.Query() 呼叫完畢後會將連線傳遞給sql.Rows型別，當然後者迭代完畢或者顯示的呼叫.Close()方法後，連線將會被釋放回到連線池(<---踩雷點，如果沒close掉是不會釋放掉連線)
	row, _ := db.Query("SELECT 1")

	PrintStats("db.Query", db)
	row.Close()
	PrintStats("row.Close()", db)

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

	//db.Begin() 呼叫完畢後將連線傳遞給sql.Tx型別物件，當.Commit()或.Rollback()方法呼叫後釋放連線。
	func() {
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
	}()

}
