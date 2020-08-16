package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// connect to bd
	_ "github.com/lib/pq"
)

// Postgres struct for working with pg db
type Postgres struct {
	db *sqlx.DB

	host     string
	port     int
	user     string
	password string
	dbName   string

	dbDriver      string
	connectLayout string
}

// Connect connect to postgres db
func (p *Postgres) Connect() (err error) {
	connecting := fmt.Sprintf(p.connectLayout, p.host, p.port, p.user, p.password, p.dbName)
	p.db, err = sqlx.Connect(p.dbDriver, connecting)
	return
}

// Ping check connection with db
func (p *Postgres) Ping() (err error) {
	err = p.db.Ping()
	return
}

// Query query to db
func (p *Postgres) Query(query string, args ...interface{}) (err error) {
	_, err = p.db.Query(query, args...)
	return
}

// Get get query to db, dest not array
func (p *Postgres) Get(dest interface{}, query string, args ...interface{}) (err error) {
	return p.db.Get(dest, query, args...)
}

// Select select query to db, dest is array
func (p *Postgres) Select(dest interface{}, query string, args ...interface{}) (err error) {
	return p.db.Select(dest, query, args...)
}

// NewPostgres construct
func NewPostgres(
	host string,
	port int,
	user string,
	password string,
	dbName string,

	dbDriver string,
	connectLayout string,
) *Postgres {
	return &Postgres{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:   dbName,

		dbDriver:      dbDriver,
		connectLayout: connectLayout,
	}
}
