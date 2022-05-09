package factory

import "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/factory_method/database_factory/db"

type PostgresqlFactory struct {
}

func (f PostgresqlFactory) GetDataBase() db.IDB {
	return new(db.PostgresqlDB)
}
