package db

import (
	"testing"

	"github.com/OlympBMSTU/exercises/db/result"
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/assert"
)

// todo - init test db
// also clear all changes

func getPool() (*pgx.ConnPool, error) {
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",   //conf.GetHost(),
			User:     "imber",       //conf.GetUser(),
			Password: "951103",      //conf.GetPassword(),
			Database: "excercieses", //conf.GetDatabase(),
		},
		MaxConnections: 2,
	}

	return pgx.NewConnPool(connPoolConfig)
}

func TestCreateSubjectOk(t *testing.T) {
	pool, err := getPool()
	if err != nil {
		t.Error("Pool not connected")
	}

	subject := "mathematics"

	expRes := result.OkResult(nil, result.CREATED)
	res := AddSubject(subject, pool)
	assert.Equal(t, expRes, res)
}

func TestCreateSubjectErrUnique(t *testing.T) {
	pool, err := getPool()
	if err != nil {
		t.Error("Pool not connected")
	}

	subject := "mathematics"

	// expRes := result.ErrorResult(result.CONSTRAINT_ERROR, "")
	AddSubject(subject, pool)
	res := AddSubject(subject, pool)
	assert.Equal(t, result.CONSTRAINT_ERROR, res.GetStatusCode())
}

func TestGetSubjectsOk(t *testing.T) {
	res := 
}

func TestGetSubjectEmptyRes(t *testing.T) {

}

func TestSaveExOk(t *testing.T) {
	// ex := entities.ExcercieseEntity{
	// 	0,
	// 	1,
	// 	"sdfd",
	// 	"dsf",
	// 	[]string{"array", "data"},
	// 	3,
	// 	"mathematic",
	// }
}

func TestSaveErrNoSubject(t *testing.T) {

}
