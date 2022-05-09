package conn

import (
	"fmt"
	"strconv"
	"sync"
)

var POOL_SIZE int = 2

var lock = &sync.Mutex{}

type connectionPool struct {
	pool []*Connection
}

var instance *connectionPool

func NewConnectionPool(poolSize int) *connectionPool {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &connectionPool{}
			instance.initPool(poolSize)
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	instance.PrintConnections()
	return instance
}

func (c *connectionPool) GetOneConnection() *Connection {
	for _, valor := range c.pool {
		if !valor.inUse {
			valor.SetInUse(true)
			return valor
		}
	}
	fmt.Println("All Connections in use !")
	return nil
}

func (c *connectionPool) initPool(poolSize int) {
	for i := 0; i < poolSize; i++ {
		conn := new(Connection)
		conn.SetId(i)
		conn.SetInUse(false)
		c.pool = append(c.pool, conn)
	}
}

func (c *connectionPool) PrintConnections() {
	for _, valor := range c.pool {
		fmt.Println("conction: " + strconv.Itoa(valor.GetId()) + " inUse: " + strconv.FormatBool(valor.IsInUse()))
	}
}
