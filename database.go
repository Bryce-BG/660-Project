package main

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Database : database wrapper
type Database struct {
	db  *sql.DB
	mux sync.Mutex
}

func Initialize() (*Database, error) {

	db, err := sql.Open("mysql", "root:@tcp(localhost)/CS660")

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS result(
		result_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		task_id INT UNSIGNED NOT NULL,
		IP TEXT NOT NULL,
		country TEXT,
		region TEXT,
		time DATETIME,
		outcome TEXT,
		user_agent TEXT,
		duration_ms FLOAT,
		PRIMARY KEY ( result_id )
	 )ENGINE=InnoDB DEFAULT CHARSET=utf8;`)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS task(
		task_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		URL TEXT NOT NULL,
		target TEXT,
		PRIMARY KEY ( task_id )
	 )ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return nil, err
	}
	return &Database{
		db: db}, err

}

func (database *Database) AddTask(url string, target string) error {
	_, err := database.db.Exec(`INSERT INTO task ( URL, target )

	 VALUES 
	 ( ?, ? );`, url, target)
	return err
}

func (database *Database) AddResultEntry(task_id int, ip, country, region string, time time.Time, user_agent string, duration float32) (int, error) {
	database.mux.Lock()
	_, err := database.db.Exec(`INSERT INTO result ( task_id, IP, country, region, time, outcome, user_agent, duration_ms  )
	 VALUES 
	 ( ?, ?, ?, ?, ?, ?, ?, ? );`, task_id, ip, country, region, time, nil, user_agent, duration)

	var id int
	if err == nil {
		id, err := database.db.Query(`SELECT LAST_INSERT_ID();`)
	}
	database.mux.Unlock()
	return id, err
}

func (database *Database) UpdateResult(task_id int, outcome string) error {

	_, err := database.db.Exec(`update result set outcome = ?`, outcome)
	return err
}

func (database *Database) PrintAllTask() error {
	res, err := database.db.Query(`select * from task`)
	defer res.Close()
	var id int
	var url, target string
	fmt.Println(res.Columns())
	for res.Next() {
		res.Scan(&id, &url, &target)
		fmt.Println(id, url, target)
	}
	return err
}

func (database *Database) Close() error {
	err := database.db.Close()
	return err
}

func main() {
	db, err := Initialize()
	if err != nil {
		return
	}
	db.AddTask("test", "test")
	db.PrintAllTask()
	db.Close()
}
