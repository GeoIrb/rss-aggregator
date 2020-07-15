package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// connect to bd
	_ "github.com/lib/pq"
)

// Postgres ...
type Postgres interface {
	Connect() (err error)
	Ping() (err error)
	Query(query string, args ...interface{}) (err error)
	Select(dest interface{}, query string, args ...interface{}) (err error)
}

type postgres struct {
	db *sqlx.DB

	host     string
	port     int
	user     string
	password string
	dbName   string

	dbDriver      string
	connectLayout string
}

func (p *postgres) Connect() (err error) {
	connecting := fmt.Sprintf(p.connectLayout, p.host, p.port, p.user, p.password, p.dbName)
	p.db, err = sqlx.Connect(p.dbDriver, connecting)
	return
}

func (p *postgres) Ping() (err error) {
	err = p.db.Ping()
	return
}

func (p *postgres) Query(query string, args ...interface{}) (err error) {
	_, err = p.db.Query(query, args...)
	return
}

func (p *postgres) Get(dest interface{}, query string, args ...interface{}) (err error) {
	return p.db.Get(dest, query, args...)
}

func (p *postgres) Select(dest interface{}, query string, args ...interface{}) (err error) {
	return p.db.Select(dest, query, args...)
}

// NewPostgres ...
func NewPostgres(
	host string,
	port int,
	user string,
	password string,
	dbName string,

	dbDriver string,
	connectLayout string,
) Postgres {
	return &postgres{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:   dbName,

		dbDriver:      dbDriver,
		connectLayout: connectLayout,
	}
}
