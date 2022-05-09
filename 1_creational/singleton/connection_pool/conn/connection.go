package conn

import (
	"fmt"
	"strconv"
)

type Connection struct {
	id    int
	inUse bool
}

func (c *Connection) Query(query string) {
	fmt.Println("Running [ " + query + " ] in Connection " + strconv.Itoa(c.id))
}

func (c *Connection) IsInUse() bool {
	return c.inUse
}

func (c *Connection) SetInUse(status bool) {
	c.inUse = status
}

func (c *Connection) GetId() int {
	return c.id
}

func (c *Connection) SetId(id int) {
	c.id = id
}
