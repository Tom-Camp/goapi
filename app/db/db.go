package db

import (
	"github.com/globalsign/mgo"
)

// DB defines the connection structure
type DB struct {
	session *mgo.Session
}

// NewConnection handles connecting to a mongo database
func NewConnection(host string, dbName string) (conn *DB) {
	session, err := mgo.Dial("mongodb://mongo:27017")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &DB{session}
	return conn
}

// Use handles connect to a certain collection
func (conn *DB) Use(dbName, tableName string) (collection *mgo.Collection) {
	return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *DB) Close() {
	conn.session.Close()
	return
}
