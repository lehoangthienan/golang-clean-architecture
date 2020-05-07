package transaction

import (
	"github.com/jinzhu/gorm"
	"github.com/lehoangthienan/marvel-heroes-backend/util/constants"
)

// TXService interface handle transaction
type TXService interface {
	TXBegin() (*Pool, error)
	TXCommit(*Pool) error
	TXRollBack(*Pool) error
}

// config struct contain config of multi service
type config struct {
	// current DB used at service
	localDB *gorm.DB
}

// NewConfig function for create new config of transaction
func NewConfig(localDB *gorm.DB) *config {
	return &config{
		localDB: localDB,
	}
}

// pool struct contain somthing to handle transaction
type Pool struct {
	// connection of current DB used of transaction
	CLocalDB *gorm.DB
	status   constants.TXStatus
}

// transaction struct
type transaction struct {
	config *config
}

// NewTransactionService func
func NewTransactionService(config *config) TXService {
	return &transaction{
		config: config,
	}
}

// TXBegin func create new transation
func (tx *transaction) TXBegin() (*Pool, error) {
	localDBTX := tx.config.localDB.Begin()
	return &Pool{
		CLocalDB: localDBTX,
		status:   constants.PendingTXStatus,
	}, nil
}

// TXCommit func commit transation
func (tx *transaction) TXCommit(pool *Pool) error {
	if pool == nil {
		return nil
	}
	if pool.status == constants.CommitedTXStatus || pool.status == constants.RollbackedTXStatus {
		return nil
	}
	var err error
	localDB := pool.CLocalDB
	if localDB != nil {
		err = localDB.Commit().Error
	}
	pool.status = constants.CommitedTXStatus
	return err
}

// TXRollBack func rollback transation
func (tx *transaction) TXRollBack(pool *Pool) error {
	if pool == nil {
		return nil
	}
	if pool.status == constants.CommitedTXStatus || pool.status == constants.RollbackedTXStatus {
		return nil
	}
	var err error
	localDB := pool.CLocalDB
	if localDB != nil {
		err = localDB.Rollback().Error
	}
	pool.status = constants.RollbackedTXStatus
	return err
}
