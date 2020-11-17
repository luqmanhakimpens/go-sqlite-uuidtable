package main

import (
    "database/sql"
    "fmt"
    "strconv"
    "time"
    //"log"       
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    database, _ := sql.Open("sqlite3", "./beacon.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS beacon (id INTEGER PRIMARY KEY, date TEXT, uuid TEXT, major INTEGER, minor INTEGER)")
    statement.Exec()
    statement, _ = database.Prepare("INSERT INTO beacon (date, uuid, major, minor) VALUES (?, ?, ?, ?)")
    
    var datetime = time.Now()
    date := fmt.Sprintf("%d-%02d-%02d", datetime.Year(), datetime.Month(), datetime.Day())

    statement.Exec(date, "40ad420f46074e77ada8d2b35f89a274", 4567, 9876) 
    
    var id int
    var date_str string
    var uuid string
    var major string
    var minor string
    
    rows, _ := database.Query("SELECT id, date, uuid, major, minor FROM beacon")
    for rows.Next() {
        rows.Scan(&id, &date_str, &uuid, &major, &minor)
        fmt.Println(strconv.Itoa(id) + " | " + date_str + " | " + uuid + " | " + major + " | " + minor)
    }

    stmt, _ := database.Prepare("select uuid FROM beacon where date = ?")
    stmt.QueryRow("2021-11-12").Scan(&uuid)
    fmt.Println(uuid)
}