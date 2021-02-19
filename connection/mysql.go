package mysqldbcon

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SetupDB() *sql.DB {
	er := godotenv.Load(".env")
	if er != nil {
		panic(er.Error())
	}
	dbDriver := os.Getenv("DB_Driver")
	dbUser := os.Getenv("DB_User")
	dbPass := os.Getenv("DB_Password")
	dbName := os.Getenv("DB_Name")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
