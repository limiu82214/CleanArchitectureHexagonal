package leveldb

import (
	"github.com/limiu82214/CleanArchitectureHexagonal/pkg/sig"
	level_db "github.com/syndtr/goleveldb/leveldb"
)

var db *level_db.DB
var err error

func GetInst() *level_db.DB {
	if db != nil {
		return db
	}
	db, err = level_db.OpenFile("../assets/leveldb/member", nil)
	if err != nil {
		sig.ShutdownServer(err)
	}
	return db
}
func DisconnectDB() {
	if db != nil {
		db.Close()
	}
}
