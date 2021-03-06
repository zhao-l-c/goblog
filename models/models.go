  package models

import (
	"os"
	"path"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
    "time"
)

const (
	_DB_NAME        = "data/beegoblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		// create directory
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		// create file
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic), new(Comment), new(Tag))
	// this one can be omitted
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func CurrentTime() string {
    return time.Now().Format("2006-01-02 15:04:05")
}

