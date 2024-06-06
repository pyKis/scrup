package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)



func ParamDb() string{
	if err := godotenv.Load(".env"); err != nil {
        log.Print("No .env file found")
    }

	host := os.Getenv("dbhost")
	port := os.Getenv("dbport")
	user := os.Getenv("dbuser")
	password := os.Getenv("dbpassword")
	name := os.Getenv("dbname")
	sslmode := os.Getenv("sslmode")


param:= "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + name + " sslmode=" + sslmode

return param
}