package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func NewMySQLConnection() (*sql.DB, error) {
	// 環境変数からデータベース接続情報を取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// 接続文字列を作成（データベース名なし）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, dbPassword, dbHost, dbPort)

	// データベース接続を確立
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 接続が確立されているか確認
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// データベースが存在しない場合は作成
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		return nil, err
	}

	// データベースを選択
	_, err = db.Exec(fmt.Sprintf("USE %s", dbName))
	if err != nil {
		return nil, err
	}

	// テーブルの作成
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id INT AUTO_INCREMENT PRIMARY KEY,
	    name VARCHAR(100) NOT NULL
	);`

	createReportsTable := `
	CREATE TABLE IF NOT EXISTS reports (
	    id INT AUTO_INCREMENT PRIMARY KEY,
	    user_id INT NOT NULL,
	    title VARCHAR(200) NOT NULL,
	    word_count INT NOT NULL,
	    writing_style VARCHAR(50) NOT NULL,
	    language VARCHAR(50) NOT NULL,
	    FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err = db.Exec(createUsersTable)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(createReportsTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
