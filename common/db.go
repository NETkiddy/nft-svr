package common

import (
	"sync"
	"time"

	"github.com/NETkiddy/common-go/config"
	"github.com/NETkiddy/common-go/log"
	svr "github.com/NETkiddy/common-go/svr_adapter/glue"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ConnectParam struct {
	Driver        string
	SourceNameKey string
}

var dbs map[string]*gorm.DB
var once sync.Once
var connectParams []*ConnectParam

func SetConnectParams(params []*ConnectParam) {
	connectParams = params
}

func InitDb() {
	once.Do(func() {
		appConfig := config.GetViper("app")

		connMaxLifeMinutes := appConfig.GetInt("mysqldata.conn_max_life_minutes")
		if connMaxLifeMinutes == 0 {
			connMaxLifeMinutes = 10
		}

		maxOpenConnections := appConfig.GetInt("mysqldata.max_open_conns")
		if maxOpenConnections == 0 {
			maxOpenConnections = 200
		}

		maxIdleConnections := appConfig.GetInt("mysqldata.max_idle_conns")
		if maxIdleConnections == 0 {
			maxIdleConnections = 100
		}

		dbs = make(map[string]*gorm.DB)

		for _, param := range connectParams {
			driver := param.Driver

			sourceName := appConfig.GetString(param.SourceNameKey)

			db, err := gorm.Open(driver, sourceName)
			if err != nil {
				log.Logger().Error(err)
				panic(&svr.ExceptionForAPIV3{Code: ErrInternalError_DBFailure, Message: "DB error"})
			}
			db.LogMode(appConfig.GetBool("gorm.log_mode"))
			db.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute)
			db.DB().SetMaxIdleConns(maxIdleConnections)
			db.DB().SetMaxOpenConns(maxOpenConnections)
			dbs[param.SourceNameKey] = db

			log.Logger().Debug("Connect database success")
		}
	})
}

func CloseDb() {
	for _, param := range connectParams {
		dbs[param.SourceNameKey].Close()

		log.Logger().Infof("Close %s success", param.SourceNameKey)
	}
}

func GetMysqlDb() *gorm.DB {
	return GetDb("mysqldata.source_name")
}

/*
获取DB读实例
*/
func GetROMysqlDb() *gorm.DB {
	return GetDb("mysqldata.ro_source_name")
}

func GetDb(sourceNameKey string) *gorm.DB {

	InitDb()

	db := dbs[sourceNameKey]

	if err := db.DB().Ping(); err != nil {
		log.Logger().Infof("ping error: %s", err.Error())
	}
	//db.SetLogger(log.Logger())

	return db
}

func GetDbBySourceName(sourceName string, driver string) *gorm.DB {

	appConfig := config.GetViper("app")

	connMaxLifeMinutes := appConfig.GetInt("mysqldata.conn_max_life_minutes")
	if connMaxLifeMinutes == 0 {
		connMaxLifeMinutes = 10
	}

	maxOpenConnections := appConfig.GetInt("mysqldata.max_open_conns")
	if maxOpenConnections == 0 {
		maxOpenConnections = 200
	}

	maxIdleConnections := appConfig.GetInt("mysqldata.max_idle_conns")
	if maxIdleConnections == 0 {
		maxIdleConnections = 100
	}

	dbs = make(map[string]*gorm.DB)

	db, err := gorm.Open(driver, sourceName)
	if err != nil {
		log.Logger().Error(err)
		panic(&svr.ExceptionForAPIV3{Code: ErrInternalError_DBFailure, Message: "DB error"})
	}
	db.LogMode(appConfig.GetBool("gorm.log_mode"))
	db.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute)
	db.DB().SetMaxIdleConns(maxIdleConnections)
	db.DB().SetMaxOpenConns(maxOpenConnections)

	//db.SetLogger(log.Logger())

	log.Logger().Debug("Connect database success")

	return db
}
