package factory

import "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/factory_method/database_factory/db"

type OracleFactory struct {
}

func (f OracleFactory) GetDataBase() db.IDB {
	return new(db.OracleDB)
}
