package databasefactory_test

import (
	"reflect"
	"testing"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/factory_method/database_factory/db"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/factory_method/database_factory/factory"
	"github.com/stretchr/testify/assert"
)

func TestDataBaseFactory(t *testing.T) {

	var dbFactory factory.IDBFactory
	var vardb db.IDB

	oracledb := new(db.OracleDB)
	postgresqldb := new(db.PostgresqlDB)

	dbFactory = new(factory.OracleFactory)

	vardb = dbFactory.GetDataBase()
	assert.Equal(t, reflect.TypeOf(oracledb), reflect.TypeOf(vardb), "is not oracle data base")

	dbFactory = new(factory.PostgresqlFactory)
	vardb = dbFactory.GetDataBase()
	assert.Equal(t, reflect.TypeOf(postgresqldb), reflect.TypeOf(vardb), "is not postgresql data base")

}
