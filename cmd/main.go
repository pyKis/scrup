package main

import (
	"log"
	p "scrup/pkg/parser"

	"scrup/db"
)






func main() {
	database,err :=db.ConnectToDB()
	if err != nil {
        log.Fatalf("Ошибка подключения к базе данных: %v\n", err)
    }
    defer 	database.Close()


	p.ParsBiggeek(database)
	
	
}