package database

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

var (
	Dbcon     *gorm.DB
	Errdb     error
	dbuser    string
	dbpass    string
	dbname    string
	dbaddres  string
	dbport    string
	dbdebug   bool
	dbtype    string
	sslmode   string
	dbtimeout string
	urlMinio  string
)

func init() {
	dbtype = beego.AppConfig.DefaultString("db.type", "")
	dbuser = beego.AppConfig.DefaultString("db.postgres.user", "")
	dbpass = beego.AppConfig.DefaultString("db.postgres.pass", "")
	dbname = beego.AppConfig.DefaultString("db.postgres.name", "")
	dbaddres = beego.AppConfig.DefaultString("db.postgres.addres", "")
	dbport = beego.AppConfig.DefaultString("db.postgres.port", "")
	sslmode = beego.AppConfig.DefaultString("db.postgres.sslmode", "")
	dbdebug = beego.AppConfig.DefaultBool("db.postgres.debug", true)
	dbtimeout = beego.AppConfig.DefaultString("db.postgres.timeout", "")

	if DbOpen() != nil {
		fmt.Println("Can Open db Postgres")
	}
}

// DbOpen ...
func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s ", dbaddres, dbport, dbuser, dbpass, dbname, sslmode, dbtimeout)
	Dbcon, Errdb = gorm.Open("postgres", args)
	fmt.Println("isi postgres sett ", args)
	if Errdb != nil {
		fmt.Println("open db Err ", Errdb)
		return Errdb
	}
	if errping := Dbcon.DB().Ping(); errping != nil {
		return errping
	}
	fmt.Println("Database connected [", dbaddres, "] [", dbname, "] [", dbuser, "] !")
	return nil
}

// GetDbCon ...
func GetDbCon() *gorm.DB {
	if errping := Dbcon.DB().Ping(); errping != nil {
		logs.Error("Db Not Connect test Ping :", errping)
		errping = nil
		if errping = DbOpen(); errping != nil {
			logs.Error("try to connect again but error :", errping)
		}
	}
	Dbcon.LogMode(dbdebug)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	Dbcon.DB().SetMaxIdleConns(200)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	Dbcon.DB().SetMaxOpenConns(200)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	Dbcon.DB().SetConnMaxLifetime(2 * time.Hour)

	return Dbcon
}
