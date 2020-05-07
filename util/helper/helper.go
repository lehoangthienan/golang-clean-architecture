package helper

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

// TransformValue func
func TransformValue(input interface{}, output interface{}) error {
	jsonObj, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonObj, output)
	return err
}

type arrayMap struct {
	repactNum  int
	atArrIndex int
}

// GetSameElementInArrays func
func GetSameElementInArrays(arrays ...[]string) []string {

	existMap := make(map[string]*arrayMap)

	for index, arr := range arrays {
		for _, value := range arr {
			if index == 0 {
				existMap[value] = &arrayMap{
					repactNum:  1,
					atArrIndex: index,
				}
			} else {
				if existMap[value] != nil && existMap[value].atArrIndex != index {
					existMap[value].repactNum++
					existMap[value].atArrIndex = index
				}
			}
		}
	}

	resultArr := []string{}

	for key := range existMap {
		if existMap[key].repactNum == len(arrays) {
			resultArr = append(resultArr, key)
		}
	}

	return resultArr
}

// DifferenceArray func
func DifferenceArray(aArr, bArr []string) []string {
	resultArr := []string{}
	diffMap := make(map[string]*string)
	for _, a := range aArr {
		diffMap[a] = &a
	}
	for _, b := range bArr {
		if diffMap[b] != nil {
			delete(diffMap, b)
		} else {
			diffMap[b] = &b
		}
	}
	for value := range diffMap {
		resultArr = append(resultArr, value)
	}
	return resultArr
}

func UseDBConn(dbConnInRepo *gorm.DB, pool *transaction.Pool) (*gorm.DB, error) {
	dbConn := dbConnInRepo
	if pool != nil && pool.CLocalDB != nil {
		dbConn = pool.CLocalDB
	}

	var err error
	if dbConn == nil {
		err = errors.NoDBConnectionError
	}

	return dbConn, err
}
