package database_test

import (
	"testing"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/singleton/database"
)

func TestDataBaseSingleton(t *testing.T) {

	d1 := database.GetInstance()
	d1.PutValue("Client 1", "Daniel")

	if !d1.HaveKey("Client 1") {
		t.Fatal("test 1: Client 1 is not in the database1")
	}

	d2 := database.GetInstance()
	d2.PutValue("Client 2", "Joana")

	if !d2.HaveKey("Client 1") {
		t.Fatal("test 2: Client 1 is not in the database2")
	}

	if !d2.HaveKey("Client 2") {
		t.Fatal("test 2: Client 2 is not in the database2")
	}

	if !d1.HaveKey("Client 2") {
		t.Fatal("test 2: Client 2 is not in the database1")
	}

	d1.DropValue("Client 1")

	if d2.HaveKey("Client 1") {
		t.Fatal("test 3: Client 1 is istill in the database1")
	}

	if d2.HaveKey("Client 1") {
		t.Fatal("test 3: Client 1 is istill in the database2")
	}
}
