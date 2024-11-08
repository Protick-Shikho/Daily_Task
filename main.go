package main

import (
    "To_do_Task/cmd"
    "To_do_Task/database"
    "fmt"
    "log"
)

func main() {


    //_____mySQL______
    // dsn := "root:123@tcp(localhost:3306)/"
    // var db database.Database = &database.MySQLDatabase{}
    

    //_____mongoDB______
    dsn := "mongodb://root:123@localhost:27017/admin"
	db := &database.MongoDB{}


	err := db.InitDB(dsn)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

    err = db.InitDB(dsn) 
    if err != nil {
        fmt.Println("Failed to initialize Database", err)
        return
    }

    cmd.SetDatabase(db)
    if err := cmd.Execute(); err != nil {
        log.Fatalf("Error executing root command: %v", err)
    }
}

