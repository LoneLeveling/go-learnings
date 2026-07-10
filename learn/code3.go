package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	id   int
	name string
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DbPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	//Adding data into database tables:
	result, err := db.Exec("insert into data values(4,'ghf')")
	checkError(err)
	lastInsertId, err := result.LastInsertId()
	fmt.Println("LastInsertId: ", lastInsertId)

	//No. of Rows affected
	rowsAffected, err := result.RowsAffected()
	fmt.Println("RowsAffected: ", rowsAffected)
	checkError(err)

	rows, err := db.Query("SELECT * FROM data")
	checkError(err)
	//NOTE: To process the rows we recived from DB as response we use a Struct
	//Below we use 2 methods Next() & scan() to loop through the rows, copy the data in our struct to print later on.
	// go doc sql Next
	for rows.Next() /*this for loop keeps running until Next() return false*/ {
		var data Data
		err := rows.Scan(&data.id, &data.name) /*note: Scan() method copies the columns from the matched row into the values that are pointed by a destination.*/
		//So in above line we have copied data from the db rows into our Struct object.
		checkError(err)
		fmt.Println(data)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
