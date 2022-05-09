package database

import "sync"

var lock = &sync.Mutex{}

type database struct {
	values map[string]string
}

var data *database

func GetInstance() *database {
	if data == nil {
		lock.Lock()
		defer lock.Unlock()
		if data == nil {
			data = &database{
				values: make(map[string]string),
			}
		}
	}

	return data
}

func (d *database) PutValue(key, value string) {
	d.values[key] = value
}

func (d *database) DropValue(key string) {
	delete(d.values, key)
}

func (d *database) HaveKey(key string) bool {
	_, present := d.values[key]

	return present
}
