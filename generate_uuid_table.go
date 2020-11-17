package main

import (
    "database/sql"     
    _ "github.com/mattn/go-sqlite3"
    "github.com/google/uuid"

    "fmt"
    "strings"
    "math/rand"
    "time"
    "os"
    "flag"
)

func main() {
    now_date_str := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
    
    db_name := flag.String("db", "beacon.db", "A name for the database")
    d_start := flag.String("start", now_date_str, "Start date, something like 2020-11-10. Default: today's system date.")
    n_of_days := flag.Int("days", 120, "Number of days")
    flag.Parse() // flag parse returns the pointer of the value

    fmt.Println("db name: ", *db_name)
    fmt.Println("start date: ", *d_start)
    fmt.Println("number of days: ", *n_of_days)
    
    os.Remove(*db_name)
    database, _ := sql.Open("sqlite3", "./" + *db_name)
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS beacon (date TEXT PRIMARY KEY, uuid TEXT, major INTEGER, minor INTEGER)")
    statement.Exec()

    random := rand.New(rand.NewSource(time.Now().UnixNano()))
    date_start, _ := time.Parse(time.RFC3339, *d_start + "T00:00:00Z")

    fmt.Println("+===============================================================+")
    fmt.Println("|   DATE     |            UUID                  | MAJOR | MINOR |")
    fmt.Println("+===============================================================+")
    defer fmt.Println("+===============================================================+") //print this at the end of the code
    
    for n :=0; n < *n_of_days; n++ {
        d := date_start.AddDate(0, 0, n)
        date := fmt.Sprintf("%d-%02d-%02d", d.Year(), d.Month(), d.Day())
        major := random.Intn(65535)
        minor := random.Intn(65535)
        uuidWithHyphen := uuid.New()
        uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

        fmt.Printf("| %s | %s | %5d | %5d |\n",date, uuid, major, minor)

        statement, _ = database.Prepare("INSERT INTO beacon (date, uuid, major, minor) VALUES (?, ?, ?, ?)")
        statement.Exec(date, uuid, major, minor) 
    }    
}