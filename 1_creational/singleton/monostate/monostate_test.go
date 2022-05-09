package monostate_test

import (
	"testing"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/singleton/monostate/conn"
	"github.com/stretchr/testify/assert"
)

func TestConnectionPoolMonostate(t *testing.T) {

	pool := conn.NewConnectionPool(2)

	conn1 := pool.GetOneConnection()
	assert.NotEmpty(t, conn1, "connection1 nil")

	conn2 := pool.GetOneConnection()
	assert.NotEmpty(t, conn2, "connection2 nil")

	conn3 := pool.GetOneConnection()
	assert.Empty(t, conn3, "connection3 not nil")

	conn1.Query("Select * from usuario")
	conn2.Query("Select * from usuario")

	pool2 := conn.NewConnectionPool(5)
	pool2.PrintConnections()

	assert.Equal(t, pool, pool2, "the pool Connections have different instances")
	conn4 := pool2.GetOneConnection()
	assert.Empty(t, conn4, "connection4 not nil")

	conn1.SetInUse(false)
	conn5 := pool2.GetOneConnection()
	assert.NotEmpty(t, conn5, "connection5 not nil")
}
