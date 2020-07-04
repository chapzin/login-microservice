package framework

import (
	"log"

	"github.com/chapzin/login-microservice/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	conn, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao db test: %v", err)
	}

	return conn
}

func NewDbProduction() *gorm.DB {
	dbInstace := NewDb()
	dbInstace.Env = "production"
	dbInstace.DbType = "mysql"
	dbInstace.Dsn = "root:123456@(localhost:3306)/micro-login?charset=utf8&parseTime=True&loc=Local"
	dbInstace.AutoMigrateDb = true
	dbInstace.Debug = true

	conn, err := dbInstace.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao db mysql: %v", err)
	}

	return conn
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.User{})
	}

	return d.Db, nil
}
