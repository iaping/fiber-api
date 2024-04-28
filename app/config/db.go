package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type Db struct {
	Enable bool   `yaml:"Enable"`
	Driver string `yaml:"Driver"`
	Mysql  Mysql  `yaml:"Mysql"`
}

func (i Db) New() (*bun.DB, error) {
	switch i.Driver {
	case "mysql":
		return i.Mysql.New()
	}

	return nil, fmt.Errorf("nonsupport: %s", i.Driver)
}

type Mysql struct {
	Host     string `yaml:"Host"`
	Port     uint16 `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
	Charset  string `yaml:"Charset"`
}

func (i Mysql) Dsn() string {
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	return fmt.Sprintf(dsn, i.User, i.Password, i.Host, i.Port, i.Database, i.Charset)
}

func (i Mysql) New() (*bun.DB, error) {
	open, err := sql.Open("mysql", i.Dsn())
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(open, mysqldialect.New())
	return db, errors.Wrap(db.Ping(), "mysql")
}
