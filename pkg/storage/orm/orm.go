package orm

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
	"strings"
	"time"
)

// Config database Config
type Config struct {
	DBType          string
	Name            string
	Addr            string
	UserName        string
	Password        string
	DisableTLS      bool
	Timezone        string
	MaxIdleConn     int
	MaxOpenConn     int
	ShowLog         bool
	LogLevel        string
	ConnMaxLifeTime time.Duration
	SlowThreshold   time.Duration // 慢查询时长，默认500ms
}

func NewOrm(c *Config) (db *gorm.DB) {
	sqlDB, err := openSql(c)
	if err != nil {
		panic(err)
	}
	switch strings.ToLower(c.DBType) {
	case "mysql":
		db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gormConfig(c))
		if err != nil {
			panic(err)
		}
		db.Set("gorm:table_options", "CHARSET=utf8mb4")
	case "postgres":
		db, err = gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), gormConfig(c))
		if err != nil {
			panic(err)
		}
	default:
		panic("database type not support")
	}
	return db
}

func openSql(c *Config) (*sql.DB, error) {
	u := url.URL{
		User: url.UserPassword(c.UserName, c.Password),
		Host: c.Addr,
		Path: c.Name,
	}

	dbType := strings.ToLower(c.DBType)

	switch dbType {
	case "mysql":
		q := make(url.Values)
		q.Set("parseTime", "true")
		q.Set("charset", "utf8mb4")
		q.Set("loc", c.Timezone)

		//u.Scheme = "mysql"
		u.RawQuery = q.Encode()
	case "postgres":
		sslMode := "require"
		if c.DisableTLS {
			sslMode = "disable"
		}

		q := make(url.Values)
		q.Set("sslmode", sslMode)
		q.Set("timezone", c.Timezone)

		u.Scheme = "postgres"
		u.RawQuery = q.Encode()

	default:
		return nil, errors.New("unsupported db type")
	}

	db, err := sql.Open(dbType, u.String())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(c.MaxIdleConn)
	db.SetMaxOpenConns(c.MaxIdleConn)
	db.SetConnMaxLifetime(c.ConnMaxLifeTime)
	return db, nil
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(c *Config) *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止外键约束, 生产环境不建议使用外键约束
		PrepareStmt:                              true, // 缓存预编译语句
	}
	return config
}
