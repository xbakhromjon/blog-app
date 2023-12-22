package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate" // install this package
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type PostgresTestSuite struct {
	suite.Suite
	db *sql.DB
	m  *migrate.Migrate
}

// Mock database name for the test
const databaseTestName = "../../app_test_db"

// Run before all tests start
func (s *PostgresTestSuite) SetupSuite() {
	// Connect mock database
	db, err := sql.Open("sqlite3", databaseTestName)
	s.db = db
	if err != nil {
		log.Fatal(err)
	}

	sqlLite3, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//Setup migration
	m, err := migrate.NewWithDatabaseInstance("file://../../migrations", databaseTestName, sqlLite3)
	s.m = m
	if err != nil {
		log.Fatal(err)
	}
	// Run migrate up
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}

// Run after each test has finished
func (s *PostgresTestSuite) TearDownTest() {
	// Truncate all tables

	tables := []string{"users", "post"}
	for _, tableName := range tables {
		deleteQuery := "delete from " + tableName
		_, err := s.db.Exec(deleteQuery)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Run after all tests has finished
func (s *PostgresTestSuite) TearDownSuite() {
	// Run migrate down
	err := s.m.Down()

	if err != nil {
		log.Fatal(err)
	}
}

func TestPostgresTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}
