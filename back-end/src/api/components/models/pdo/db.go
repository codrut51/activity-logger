package pdo

import (
	"database/sql"
	"sync"
	_ "github.com/lib/pq"
	"api/components/models/config"
)

var lock = &sync.Mutex{}

var db *sql.DB 

func GetInstance() *sql.DB {
	lock.Lock()
	defer lock.Unlock()

	if db == nil {
		cfg := config.GetConfig()
		var err error
		dbCfg = cfg.Database;
		psqlInfo := fmt.Spritf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
							   dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.Name )
		dbOne, err = sql.Open("postgres", psqlInfo) 		
		if err != nil {
			return db
		}
		db = make(dbOne)
	}
	return db
}