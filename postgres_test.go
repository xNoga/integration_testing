package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	user   = "appdev"
	dbname = "appdev"
)

func setupDB(user, dbname string) (db *sql.DB, err error) {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)
	db, err = sql.Open("postgres", dbinfo)
	return db, err
}

func TestDB(t *testing.T) {
	db, err := setupDB(user, dbname)
	assert.Nil(t, err)

	err = db.Ping()
	assert.Nil(t, err)
}

func TestGetTracks(t *testing.T) {
	db, err := setupDB(user, dbname)
	assert.Nil(t, err)

	res, err := getTracks(db)
	assert.Nil(t, err)

	assert.Equal(t, 39, len(res), "The length of the result set should be 39")
}

func TestGetInvoices(t *testing.T) {
	db, err := setupDB(user, dbname)
	assert.Nil(t, err)

	res, err := getInvoices(db)
	assert.Nil(t, err)

	assert.Equal(t, 115, len(res), "The length of the result set should be 115")
}

func TestGetCustomers(t *testing.T) {
	db, err := setupDB(user, dbname)
	assert.Nil(t, err)

	res, err := getCustomers(db)
	assert.Nil(t, err)

	assert.Equal(t, 11, len(res), "The length of the result set should be 11")
}

func TestGetAlbums(t *testing.T) {
	db, err := setupDB(user, dbname)
	assert.Nil(t, err)

	res, err := getAlbums(db)
	assert.Nil(t, err)

	assert.Equal(t, 12, len(res), "The length of the result set should be 12")
}
