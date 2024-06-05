package db

import (
	"database/sql"

	"log"

	c "scrup/configs"

	_ "github.com/lib/pq"
)



func ConnectToDB() {

psqlInfo := c.ParamDb()

db, err := sql.Open("postgres", psqlInfo)
if err != nil {
    log.Fatalf("Невозможно подключиться к базе данных: %v\n", err)
}
defer db.Close()

err = db.Ping()
if err != nil {
    log.Fatalf("Не удалось получить доступ к базе данных: %v\n", err)
}

log.Println("Успешно подключено к базе данных!")
}