package db

type IDB interface {
	Query(sql string)
	Update(sql string)
}
