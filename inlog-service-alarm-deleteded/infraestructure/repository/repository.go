package repository

import (
	"log"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Repository db connectionn
type Repository struct {
	Conn *sqlx.DB
	Tx   *sqlx.Tx
}

// EventHandler call execute in transaction
type EventHandler func() error

var _db string
var _connectionString string

// SetDb postgres, firebrid, mysql ...
func setDb(db, connectionString string) {
	_db = db
	_connectionString = connectionString
}

func open() *sqlx.DB {
	db, err := sqlx.Connect(_db, _connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// ExecInTransaction executa conjunto de comandos em transacao
func (r Repository) ExecInTransaction(eh EventHandler) error {
	if r.Conn == nil {
		r.Conn = open()
	}
	defer r.Conn.Close()
	r.Tx = r.Conn.MustBegin()
	err := eh()
	if err != nil {
		r.Tx.Rollback()
		return err
	}
	err = r.Tx.Commit()
	r.Tx = nil
	return err
}

// Query return select query db
func (r Repository) Query(rs interface{}, query string, param ...interface{}) error {
	str := reflect.TypeOf(rs).String()
	if strings.Contains(str, "[]") {
		return getMany(rs, query, param...)
	}
	return getOne(rs, query, param...)
}

// GetOne return select query db multiples rows
func getMany(rs interface{}, query string, param ...interface{}) error {
	conn := open()
	defer conn.Close()
	if len(param) == 0 {
		return conn.Select(rs, query)
	}
	rows, _ := conn.NamedQuery(query, param[0])
	return sqlx.StructScan(rows, rs)
}

// GetOne return select query db single rows
func getOne(rs interface{}, query string, param ...interface{}) (er error) {
	conn := open()
	defer conn.Close()
	if param == nil {
		return conn.Get(rs, query)
	}

	rows, _ := conn.NamedQuery(query, param[0])
	for rows.Next() {
		er = rows.StructScan(rs)
	}
	return
}

// Exec exec sql insert update delete
func (r Repository) Exec(query string, param interface{}, err interface{}) (rs interface{}) {
	if err != nil {
		return
	}
	var errorExec error
	var rows *sqlx.Rows
	if r.Tx == nil {
		errorExec = &errorString{"execute query, INSERT, UPDATE OR DELETE in transaction!"}
		return nil
	}
	rows, errorExec = r.Tx.NamedQuery(query, param)
	if rows.Next() {
		errorExec = rows.Scan(&rs)
	}
	if err == nil && errorExec != nil {
		err = errorExec
	}
	return
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
