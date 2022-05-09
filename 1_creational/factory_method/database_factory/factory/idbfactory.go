package factory

import "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/factory_method/database_factory/db"

type IDBFactory interface {
	GetDataBase() db.IDB
}
