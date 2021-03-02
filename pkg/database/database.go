package database

import (
	"database/sql"
	"time"
	"goblog/pkg/logger"

	"github.com/go-sql-driver/mysql"
)

// DB 数据库对象
var DB *sql.DB

//  Initialize 初始化数据库
func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(151)
	// 设置最大空间连接数
	DB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)
	// 尝试链接, 失败会报错
	err = DB.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
	id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	body longtext COLLATE utf8mb4_unicode_ci
);`
	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
