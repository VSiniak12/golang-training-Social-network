package data

import (
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var testState = &State{
	IdState:   17,
	Name:      "Gomelskaia",
	CountryId: 2,
}

func TestStateData_ReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)
	rows := sqlmock.NewRows([]string{"id_state", "name", "country_id"}).
		AddRow(testState.IdState, testState.Name, testState.CountryId)
	mock.ExpectQuery(readAllStatesQuery).WillReturnRows(rows)
	states, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(states)
	assert.Equal(states[0], *testState)
	assert.Len(states, 1)
}

func TestStateData_ReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)
	mock.ExpectQuery(readAllStatesQuery).WillReturnError(errors.New("something went wrong in func ReadAll"))
	states, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(states)
}

func TestStateData_Read(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)
	rows := sqlmock.NewRows([]string{"id_state", "name", "country_id"}).
		AddRow(testState.IdState, testState.Name, testState.CountryId)
	mock.ExpectQuery(readStateQuery).WillReturnRows(rows)
	state, err := data.Read(17)
	assert.NoError(err)
	assert.NotEmpty(state)
	assert.Equal(state, testState)
}

func TestStateData_ReadErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)
	mock.ExpectQuery(readStateQuery).WillReturnError(errors.New("something went wrong in func Read"))
	state, err := data.Read(17)
	assert.Error(err)
	assert.Empty(state)
}

func TestStateData_Add(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)
	mock.ExpectBegin()
	mock.ExpectExec(insertStateQuery).
		WithArgs(testState.IdState, testState.Name, testState.CountryId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	id, err := data.Add(testState)
	assert.NoError(err)
	assert.NotEmpty(id)
	assert.Equal(id, testState.IdState)
}

func TestAddErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)
	mock.ExpectBegin()
	mock.ExpectExec(insertStateQuery).
		WithArgs(testState.IdState, testState.Name, testState.CountryId).
		WillReturnError(errors.New("something went wrong in func Add"))
	mock.ExpectCommit()
	id, err := data.Add(testState)
	assert.Error(err)
	assert.Equal(-1, id)
}

func TestStateData_Update(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)
	mock.ExpectBegin()
	mock.ExpectExec(insertStateQuery).
		WithArgs(testState.IdState, testState.Name, testState.CountryId).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	state, err := data.Update(testState)
	assert.NoError(err)
	assert.NotEqual(nil, state)
}

func TestUpdateErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}

	data := NewStateData(dbGorm)
	mock.ExpectQuery(insertStateQuery).WithArgs(testState.IdState).
		WillReturnError(errors.New("something went wrong in func Update"))
	_, err = data.Update(testState)
	assert.Error(err)
}

func TestStateData_Delete(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)

	rows := mock.NewRows([]string{"id_state", "name", "country_id"}).
		AddRow(testState.IdState, testState.Name, testState.CountryId)
	mock.ExpectQuery(selectStatesWhereQuery).WithArgs(testState.IdState).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(deleteStateQuery).WithArgs(testState.IdState).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = data.Delete(17)
	assert.NoError(err)
	err = data.Delete(0)
	assert.Error(err)
}

func TestDeleteErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewStateData(dbGorm)
	rows := mock.NewRows([]string{"id_state", "name", "country_id"}).
		AddRow(testState.IdState, testState.Name, testState.CountryId)
	mock.ExpectQuery(selectStatesWhereQuery).WithArgs(testState.IdState).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(deleteStateQuery).WithArgs(testState.IdState).
		WillReturnError(errors.New("something went wrong in func Delete"))
	mock.ExpectCommit()
	err = data.Delete(5)
	assert.Error(err)
}
