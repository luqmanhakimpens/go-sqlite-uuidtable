package main

import (
    "database/sql"
    "fmt"
    "time"
    _ "github.com/mattn/go-sqlite3"
    "flag"
)

func main() {
    now_date_str := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
    
    date := flag.String("date", now_date_str, "Date, something like 2020-11-10. Default: today's system date. ")
    db_name := flag.String("db", "beacon.db", "The name of the database")
    flag.Parse() // flag parse returns the pointer of the value

    database, _ := sql.Open("sqlite3", "./" + *db_name)
        
    var uuid string
    var major string
    var minor string
    
    stmt, _ := database.Prepare("select uuid, major, minor FROM beacon WHERE date = ?")
    stmt.QueryRow(*date).Scan(&uuid, &major, &minor)
    fmt.Println(*date + " | " + uuid + " | " + major + " | " + minor)
}