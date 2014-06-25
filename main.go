package main

import (
    "database/sql"
    "fmt"
    "github.com/mattn/go-sqlite3"
    "log"
    "os"
)

func main() {
    //Remove current db file
    os.Remove("./sqlite.db")

    db, err := sql.Open("sqlite3", "./sqlite.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    sql := `
		CREATE TABLE users(username PRIMARY KEY, firstlast, password);
	`
    _, err = db.Exec(sql)
    if err != nil {
        log.Printf("%q: %s\n", err, sql)
        return
    }

    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }

    stmt, err := tx.Prepare("insert into users(username, firstlast, password) values(?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }

    defer stmt.Close()

    _, err = stmt.Exec("admin", "Nick Robinson", "password")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf(">")
    for {
        var str, realname, username, password string
        fmt.Scanf("%s", &str)
        switch str {
        default:
            fmt.Printf("echo %s\n", str)
            fmt.Printf(">")
        case "exit":
            return
        case "enable":
            fmt.Printf("Username: ")
            fmt.Scanf("%s", &username)
            fmt.Printf("Password: ")
            fmt.Scanf("%s", &password)
            err := db.QueryRow("SELECT b FROM users WHERE a=? AND c=?", username, password).Scan(&realname)
            if err != nil {
                log.Fatal(err)
            }
            fmt.Println(realname)
        }
    }
}
