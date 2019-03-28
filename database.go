package main

import (
	"database/sql"
	"fmt"
	"log"
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

func (database *Database) OfferRandomTask() (int, string, error) {
	res, err := database.db.Query(`select * from task order by rand() limit 1`)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	var id int
	var url, target string
	fmt.Println(res.Columns())
	for res.Next() {
		res.Scan(&id, &url, &target)
		fmt.Println(id, url, target)
	}
	return id, url, err
}

func (database *Database) AddResultEntry(task_id int, ip, user_agent string) (int, error) {
	database.mux.Lock()
	_, err := database.db.Exec(`INSERT INTO result ( task_id, IP, user_agent)
	 VALUES 
	 ( ?, ?, ?);`, task_id, ip, user_agent)

	var id int
	if err == nil {
		res, err := database.db.Query(`SELECT LAST_INSERT_ID();`)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Close()
		for res.Next() {
			res.Scan(&id)
		}
	}
	database.mux.Unlock()
	fmt.Println(id)
	return id, err
}

func (database *Database) AddResul(task_id int, ip, country, region string, time time.Time, user_agent string, duration float32) (int, error) {
	database.mux.Lock()
	_, err := database.db.Exec(`INSERT INTO result ( task_id, IP, country, region, time, outcome, user_agent, duration_ms  )
	 VALUES 
	 ( ?, ?, ?, ?, ?, ?, ?, ? );`, task_id, ip, country, region, time, nil, user_agent, duration)

	var id int
	if err == nil {
		res, err := database.db.Query(`SELECT LAST_INSERT_ID();`)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Close()
		for res.Next() {
			res.Scan(&id)
		}
	}
	database.mux.Unlock()
	fmt.Println(id)
	return id, err
}

func (database *Database) UpdateResult(result_id int, outcome string) error {

	_, err := database.db.Exec(`update result set outcome = ? where result_id = ?`, outcome, result_id)
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

// func main() {
// 	db, err := Initialize()
// 	if err != nil {
// 		return
// 	}
// 	db.AddTask("https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_92x30dp.png", "google")
// 	db.AddTask("https://en.wikipedia.org/static/images/project-logos/enwiki.png", "wikipedia ")
// 	db.PrintAllTask()
// 	db.Close()
// }
