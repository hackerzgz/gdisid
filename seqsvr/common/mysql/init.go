package mysql

import (
	"bytes"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB     *sql.DB
	dbOnce *sync.Once
)

// Config databases
type Config struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	DBName      string `json:"db_name"`
	Params      string `json:"params"`
	MaxLifetime int    `json:"max_lifetime"`
	PoolSize    int    `json:"pool_size"`
}

// DSN return data source name to connect mysql
func (s Config) DSN() string {
	var buf bytes.Buffer

	if s.User != "" {
		buf.WriteString(s.User)
		if s.Password != "" {
			buf.WriteByte(':')
			buf.WriteString(s.Password)
		}

		buf.WriteByte('@')
	}
	if s.Host != "" {
		buf.WriteString("tcp(")
		buf.WriteString(s.Host)
		if s.Port > 0 {
			buf.WriteByte(':')
			buf.WriteString(fmt.Sprintf("%d", s.Port))
		}
		buf.WriteByte(')')
	}
	buf.WriteByte('/')
	buf.WriteString(s.DBName)

	if s.Params != "" {
		buf.WriteByte('?')
		buf.WriteString(s.Params)
	} else {
		buf.WriteString("?charset=utf8mb4&autocommit=true&parseTime=True")
	}

	return buf.String()
}

// New init databases connections and panic if failed to connect
func New(config Config) error {
	dbOnce.Do(func() {
		db, err := sql.Open("mysql", config.DSN())
		if err != nil {
			panic(fmt.Errorf("failed to open db: %v", err))
		}

		err = db.Ping()
		if err != nil {
			panic(fmt.Errorf("failed to ping db: %v", err))
		}

		db.SetConnMaxLifetime(time.Duration(4*60*60) * time.Second)
		if config.MaxLifetime > 0 {
			db.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
		}

		psize := func() int {
			if config.PoolSize > 0 {
				return int(config.PoolSize)
			}

			return 100
		}()
		db.SetMaxIdleConns(psize)
		db.SetMaxOpenConns(psize)
	})

	return nil
}
