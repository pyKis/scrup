package db

import (
	"database/sql"
	"fmt"

	"log"

	c "scrup/configs"
	"scrup/models"

	_ "github.com/lib/pq"
)

var db *sql.DB





func ConnectToDB()(*sql.DB, error) {

psqlInfo := c.ParamDb()

db, err := sql.Open("postgres", psqlInfo)
if err != nil {
    log.Fatalf("Невозможно подключиться к базе данных: %v\n", err)
}

log.Println("Соединение с базой данных открыто")
err = db.Ping()
if err != nil {
	
    log.Fatalf("Не удалось получить доступ к базе данных: %v\n", err)
}

log.Println("Подключено к базе данных!")

//createTable(db)


log.Println("База данных создана!")

return db, nil
}

func createTable(db *sql.DB) {
createTableSQL := "CREATE TABLE items (name VARCHAR(250) NOT NULL, linc VARCHAR(250) NOT NULL, price VARCHAR(20) NOT NULL, oldPrice VARCHAR(20) NOT NULL)"


_, err := db.Exec(createTableSQL)
    if err != nil {
        log.Fatalf("Не удается создать таблицу: %v\n", err)
    }

    log.Println("Таблица создана!")
}


func InsertItemBiggeek(db *sql.DB,item models.Item) error {
	if db == nil {
        return fmt.Errorf("db is nil")
    }

	query := `INSERT INTO items (name, linc, price, oldPrice) VALUES ($1, $2, $3, $4)`
    _, err := db.Exec(query, item.Name, item.Linc, item.Price, item.OldPrice)
	if err != nil {
		log.Fatalf("Не удается вставить данные: %v\n", err)		
	}


    return  nil
}