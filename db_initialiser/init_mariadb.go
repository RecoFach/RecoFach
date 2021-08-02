package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	//"github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//struct to perform operations on the MariaDB
type Database struct{}

//reads the csv file with subjects data, connects to Maria and loads the csv data into mariaDB
func main() {
	rows := readOrders("head.csv")
	mariaDb := Database{}
	mariaDb.insertRowsToDatabase(rows)
}

//reading csv data with subjects data and storing it as an array of strings
func readOrders(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}

//creating the schema to insert furhter the subjects data
func (mariaDb *Database) createSubjectsSchema(db *sql.DB) {
	_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS Curriculum")
	if err != nil {
		fmt.Println("schema")
		panic(err)
	}
}

func (mariaDb *Database) createSubjectsTable(db *sql.DB) {
	_, err := db.Exec("Create TABLE IF NOT EXISTS Curriculum.Subjects (    CourseName varchar(255),    Link varchar(255),    Sws int ,german varchar(255)	, english varchar(255) ,Komplexpraktikum int, Seminar int , Vorlesung int, SoftwareEngineering int, AI int, LowLevel int, Security int, Web int, Theoretical int,Sommersemester int, Wintersemester int)")
	if err != nil {
		print("table")
		panic(err)
	}
}

//insert the csv data into MariaDB
func (mariaDb *Database) insertRowsToDatabase(rows [][]string) {
	// Connect to the database.
	//connString := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	//connecting to the mariadb server without specifying a specific DB schema
	connString := "root:root@tcp(maria:3306)/"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	mariaDb.createSubjectsSchema(db)
	mariaDb.createSubjectsTable(db)

	//Insert the rows, omitting the first header row from the CSV.
	stmt, err := db.Prepare("INSERT INTO Curriculum.Subjects VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range rows[1:] {
		_, err := stmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14], row[15])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Successfull")

}

func db_connect() {
	log.Println("Start Software")
	dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dsn)
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
	insert, err := db.Query("INSERT INTO test.table VALUES ( 'A', 2002 )")
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
