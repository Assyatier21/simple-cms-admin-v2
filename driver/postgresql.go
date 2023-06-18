package driver

import (
	"database/sql"
	"log"

	"github.com/assyatier21/simple-cms-admin-v2/config"
)

func InitPostgres(cfg config.DBConfig) *sql.DB {
	psqlInfo := cfg.GetDSN()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("[Database] initialized...")

	err = db.Ping()
	if err != nil {
		log.Println("[Database] failed to connect to database: ", err)
		return nil
	}

	log.Println("[Database] successfully connected")
	return db
}
