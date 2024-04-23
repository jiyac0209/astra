package repository

import (
	"astraSecurity/domain"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "postgres"
)

func Connect() *sql.DB {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")
	return db
}

func CreateSchema(db *sql.DB) error {
	fmt.Println("Creating schema...")
	// Define the SQL statement to create the table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS Astra (
		uuid varchar(50) PRIMARY KEY,
		data VARCHAR(100),
		Timestamp timestamp
	);
	`

	// Execute the SQL statement to create the table
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	fmt.Println("Created schema")

	return nil
}

func Insert(db *sql.DB, astra domain.Astra) error {

	// Define the SQL statement to insert data into the database
	insertSQL := `
		INSERT INTO Astra (uuid, data, timestamp)
		VALUES ($1, $2, $3)
	`

	// Execute the SQL statement to insert data into the database
	_, err := db.Exec(insertSQL, astra.UUID, astra.Data, astra.Timestamp)
	if err != nil {
		return err
	}

	return nil
}
