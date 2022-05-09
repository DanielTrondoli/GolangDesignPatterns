package db

import "fmt"

type OracleDB struct{}

func (db OracleDB) Query(sql string) {
	fmt.Println("Querying " + sql + " in Oracle database")
}

func (db OracleDB) Update(sql string) {
	fmt.Println("Uptade querry " + sql + " in Oracle database")
}
