package models

import (
	"github.com/hiromaily/golibs/db/mysql"
)

//
//extension of db/mysql package.
//

// Models is extension of mysql.DBInfo
type Models struct {
	Db *mysql.MS
}

var db Models

// when making mysql instance, first you should use mysql.New()
func new() {
	db = Models{}
	db.Db = mysql.GetDB()
}

// GetDB is to get mysql instance. it's using singleton design pattern.
func GetDB() *Models {
	if db.Db == nil {
		new()
	}
	return &db
}
