package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//                            user password     server or ip        database
	db, err := sql.Open("mysql", "root:rootroot@tcp(172.31.212.242:3306)/servddb")

	if err != nil {
		panic(err)
	}
	db.Query("INSERT INTO user VALUES ( 'najm' );")
	db.Query("drop table user;")
	db.Query("create table user(age INT NOT NULL, email VARCHAR(40) NOT NULL, name VARCHAR(20) NOT NULL, PRIMARY KEY (email));")

	db.Query("INSERT INTO user VALUES ( 20 , 'najm@najm.com', 'najm' );")
	db.Query("INSERT INTO user VALUES ( 30 , 'najm1@najm.com', 'Najm' );")
	db.Query("INSERT INTO user VALUES ( 40 , 'najm2@najm.com', 'Najm' );")
	db.Query("INSERT INTO user VALUES ( 20 , 'najm3@najm.com', 'najm' );")
	fmt.Printf("inserted\n")

	result, err := db.Query("SELECT * FROM user;")
	// handle error
	if err != nil {
		panic(err)
	}

	for result.Next() {

		var age int
		var email string
		var name string
		err = result.Scan(&age, &email, &name)

		// handle error
		if err != nil {
			panic(err)
		}
		fmt.Printf("Age: %d Email: %s Name: %s\n", age, email, name)

	}

	db.Query("ALTER TABLE user DROP COLUMN name;")
	fmt.Printf("ALTERED!!\n")
	rlt, err := db.Query("SELECT * FROM user;")
	// handle error
	if err != nil {
		panic(err)
	}

	for rlt.Next() {

		var age int
		var email string
		err = rlt.Scan(&age, &email)

		// handle error
		if err != nil {
			panic(err)
		}
		fmt.Printf("Age: %d Email: %s \n", age, email)

	}

	db.Query("DELETE  FROM user WHERE age = 20; ")
	fmt.Printf("deleted\n")

	res, err := db.Query("SELECT * FROM user;")
	// handle error
	if err != nil {
		panic(err)
	}

	// th.
	for res.Next() {

		var age int

		var email string

		err = res.Scan(&age, &email)

		// handle error
		if err != nil {
			panic(err)
		}
		fmt.Printf("Age: %d Email: %s\n", age, email)

	}
	db.Query("DELETE  FROM user WHERE age = 40; ")

	defer db.Close()

	fmt.Printf("ended\n")

}
