package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"time"
)

var host, user, password, database, table string

func uploadFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	now := time.Now()
	nowFormat := now.Format("2006-01-02 15:04:05")

	urlInsert, err := db.Prepare(fmt.Sprintf(`INSERT INTO %s (name, comment, logo, create_time, update_time) VALUES ( ?, ?, ?, ?, ? )`, table))
	if err != nil {
		log.Fatal(err)
	}
	defer urlInsert.Close()

	_, err = urlInsert.Exec(fileName, "", data, nowFormat, nowFormat)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("insert file %s into table %s done.", fileName, table))
}

func main() {
	var fileName string
	files := []string{}

	flag.StringVarP(&host, "host", "h", "localhost", "specify mysql host, default localhost")
	flag.StringVarP(&user, "user", "u", "root", "specify mysql user, default root")
	flag.StringVarP(&password, "password", "p", "", "specify mysql password, default empty")
	flag.StringVarP(&database, "database", "D", "", "specify mysql database")
	flag.StringVarP(&table, "table", "t", "", "specify mysql table")

	flag.StringVarP(&fileName, "file", "f", "", "specify one file")
	//flag.StringVar(&files, "files", "", "specify all files")
	//flag.StringArrayVar(&files, "files", nil, "specify all files")
	flag.StringSliceVarP(&files, "files", "F", nil, "specify more files")

	flag.Parse()

	fmt.Println(files)

	if fileName != "" {
		uploadFile(fileName)
	}

	if files != nil {
		for _, f := range files {
			fmt.Println(f)
			uploadFile(f)
		}
	}
}
