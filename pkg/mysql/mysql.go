package mysql

import (
	"fmt"
	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	Database string `json:"database" mapstructure:"database"`
}

type DB struct {
	DB *gorm.DB
}

func New(cfg Config) *DB {
	dsnConfig := NewDBConfig(cfg)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsnConfig.FormatDSN(),
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when init mysql db connection: %s", err)
	}

	return &DB{DB: db}
}

func NewDBConfig(cfg Config) *mysqldriver.Config {
	//loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	dsnConfig := mysqldriver.NewConfig()
	dsnConfig.User = cfg.Username
	dsnConfig.Passwd = cfg.Password
	dsnConfig.Net = "tcp"
	dsnConfig.Addr = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	dsnConfig.DBName = cfg.Database
	dsnConfig.Params = map[string]string{"charset": "utf8mb4"}
	dsnConfig.ParseTime = true
	dsnConfig.AllowNativePasswords = true
	//dsnConfig.Loc = Local

	return dsnConfig
}
