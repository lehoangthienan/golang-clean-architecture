package transaction

import (
	"testing"

	pgConfig "github.com/lehoangthienan/marvel-heroes-backend/util/config/db/pg"
)

// Test_tx_Begin func create new transation
func Test_tx_Begin(t *testing.T) {
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	tx := &transaction{
		&config{
			localDB: dbTest,
		},
	}

	err := tx.config.localDB.Begin().Error

	if err != nil {
		t.Fatalf("Test_tx_Begin failed by error: %v", err)
	}
}

// Test_tx_Commit func commit transation
func Test_tx_Commit(t *testing.T) {
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	tx := &transaction{
		&config{
			localDB: dbTest,
		},
	}

	localDBTX := tx.config.localDB.Begin()

	err := localDBTX.Commit().Error

	if err != nil {
		t.Fatalf("Test_tx_Commit failed by error: %v", err)
	}
}

// Test_tx_RollBack func rollback transation
func Test_tx_RollBack(t *testing.T) {
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	tx := &transaction{
		&config{
			localDB: dbTest,
		},
	}

	localDBTX := tx.config.localDB.Begin()

	err := localDBTX.Rollback().Error

	if err != nil {
		t.Fatalf("Test_tx_RollBack failed by error: %v", err)
	}
}
