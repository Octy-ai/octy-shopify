package database

import (
	"database/sql"
	"log"

	"github.com/Octy-ai/octy-shopify/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

type Adapter struct {
	conn   *sql.DB
	config *config.Config
}

func NewAdapter(config *config.Config) (*Adapter, error) {
	conn := &sql.DB{}
	return &Adapter{conn: conn, config: config}, nil
}

func (dba *Adapter) Connect() {

	log.Printf("Connecting to database at: %v ... \n", dba.config.App.DBPath)
	connection, err := sql.Open("sqlite3", dba.config.App.DBPath)
	if err != nil {
		log.Fatalln("Failed to connect to Database: ", err)
	}
	(*dba).conn = connection
	log.Printf("Connected to database at: %v !\n", dba.config.App.DBPath)

	// Create customer table
	statement, err :=
		dba.conn.Prepare("CREATE TABLE IF NOT EXISTS Customers (ID INTEGER PRIMARY KEY AUTOINCREMENT, octy_customer_id VARCHAR(255) NOT NULL, octy_profile_id VARCHAR(255) NOT NULL, shopify_customer_id VARCHAR(255) NULL, CreatedAt DATETIME NOT NULL, UpdatedAt DATETIME NULL)")
	if err != nil {
		log.Fatalln("Failed to connect to Database: ", err)
	}
	statement.Exec()
}
