package main

import (
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	h := md5.New()
	defer db.Close()
	db.Exec("CREATE TABLE users(a PRIMARY KEY, b, c)")
	io.WriteString(h, "password")
	args := sqlite3.NamedArgs{"$a": "admin", "$b": "Nick Robinson", "$c": h.Sum(nil)}
	db.Exec("REPLACE INTO users VALUES($a, $b, $c)", args)
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
			switch {
			case err == sql.ErrNoRows:
				log.Printf("No user with that ID.")
				fmt.Printf(">")
			case err != nil:
				log.Fatal(err)
			default:
				fmt.Printf("Username is %s\n", realname)
				fmt.Printf(">")
			}
		}
	}
}
