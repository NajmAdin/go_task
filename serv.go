package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"net"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:rootroot@tcp(0.0.0.0:3306)/servddb")

	if err != nil {
		panic(err)
	}

	fmt.Println("start")

	PORT := ":5020"
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		db.Query("INSERT INTO user VALUES('" + string(netData) + "');")

		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
	defer db.Close()
}

/*


        	db, err := sql.Open("mysql", "root:rootroot@tcp(0.0.0.0:3306)/taskone")
if err != nil {
		panic(err)
	}
	db.Query("DELETE  FROM user WHERE age = 40; ")

	defer db.Close()
*/
