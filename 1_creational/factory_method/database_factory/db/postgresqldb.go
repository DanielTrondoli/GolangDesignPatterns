package db

import "fmt"

type PostgresqlDB struct{}

func (db PostgresqlDB) Query(sql string) {
	fmt.Println("Querying " + sql + " in Oracle database")
}

func (db PostgresqlDB) Update(sql string) {
	fmt.Println("Uptade querry " + sql + " in Oracle database")
}
