package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type track struct {
	trackid      int
	name         string
	albumid      int
	mediatypeid  int
	genreid      int
	composer     sql.NullString
	milliseconds int64
	bytes        int64
	unitprice    float64
}

type invoice struct {
	invoiceid      int
	customerid     int
	invoicedate    string
	billingaddress string
}

type customer struct {
	customerid int
	firstname  string
	lastname   string
	company    sql.NullString
}

type album struct {
	albumid  int
	title    string
	artistid int
	composer string
}

func getTracks(db *sql.DB) (res []*track, err error) {
	rows, err := db.Query(`
		(SELECT * FROM chinook.track WHERE genreid = $1) 
		UNION 
		(SELECT * FROM chinook.track WHERE genreid = $2) 
		ORDER BY trackid;`, 18, 20)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		t := &track{}
		err := rows.Scan(&t.trackid, &t.name, &t.albumid, &t.mediatypeid, &t.genreid, &t.composer, &t.milliseconds, &t.bytes, &t.unitprice)
		if err != nil {
			panic(err)
		}
		res = append(res, t)
	}

	return
}

func getInvoices(db *sql.DB) (res []*invoice, err error) {
	rows, err := db.Query(`
		(SELECT invoiceid, customerid, invoicedate, billingaddress FROM chinook.invoice WHERE total < $1) 
		INTERSECT 
		(SELECT invoiceid, customerid, invoicedate, billingaddress FROM chinook.invoice WHERE total > $2) 
		ORDER BY invoiceid`, 10, 5)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		i := &invoice{}
		err := rows.Scan(&i.invoiceid, &i.customerid, &i.invoicedate, &i.billingaddress)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}

	return
}

func getCustomers(db *sql.DB) (res []*customer, err error) {
	rows, err := db.Query(`
		(SELECT customerid, firstname, lastname, company FROM chinook.customer WHERE country = 'USA') 
		EXCEPT 
		(SELECT customerid, firstname, lastname, company FROM chinook.customer where email LIKE '%yahoo.com')
		`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		c := &customer{}
		err := rows.Scan(&c.customerid, &c.firstname, &c.lastname, &c.company)
		if err != nil {
			panic(err)
		}
		res = append(res, c)
	}

	return
}

func getAlbums(db *sql.DB) (res []*album, err error) {
	rows, err := db.Query(`
		(SELECT chinook.album.*, chinook.track.composer
		FROM chinook.album, chinook.track
		WHERE album.albumid = track.albumid
		AND track.composer LIKE '%Mozart')
		UNION
		(SELECT chinook.album.*, chinook.track.composer
		FROM chinook.album, chinook.track
		WHERE album.albumid = track.albumid
		AND track.composer LIKE '%Bach')
		ORDER BY composer
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		a := &album{}
		err := rows.Scan(&a.albumid, &a.title, &a.artistid, &a.composer)
		if err != nil {
			panic(err)
		}
		res = append(res, a)
	}

	return
}
