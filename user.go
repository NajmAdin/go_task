package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
)

func main() {
        db, err := sql.Open("mysql", "root:rootroot@tcp(0.0.0.0:3306)/userddb")

        fmt.Println("start")

        if err != nil {
		panic(err)
	}
        

        CONNECT := "127.0.0.1:5020"
        c, err := net.Dial("tcp", CONNECT)
        if err != nil {
                fmt.Println(err)
                return
        }

        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")
                text, _ := reader.ReadString('\n')
                fmt.Fprintf(c, text+"\n")
                db.Query("INSERT INTO user VALUES('" + text + "');")

                message, _ := bufio.NewReader(c).ReadString('\n')
                fmt.Print("->: " + message)
                if strings.TrimSpace(string(text)) == "STOP" {
                        fmt.Println("TCP client exiting...")
                        return
                }
        }
        defer db.Close()
}
    
