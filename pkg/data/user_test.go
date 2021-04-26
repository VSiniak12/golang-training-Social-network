package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

var testUser = &User{
	IdUser:    3,
	Login:     "MockLogin",
	Password:  "MockPassword",
	Gender:    true,
	Email:     "mock@mail.ru",
	LastName:  "MockLastName",
	FirstName: "MockFirstName",
	Birthday:  time.Date(2000, 3, 10, 0, 0, 0, 0, time.UTC),
	CityId:    2,
}

func TestUserData_ReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	rows := sqlmock.
		NewRows([]string{"id_user", "login", "password", "gender", "email",
			"last_name", "first_name", "birthday", "city_id"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityId)
	mock.ExpectQuery(readAllUsersQuery).WillReturnRows(rows)
	users, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(users)
	assert.Equal(users[0], *testUser)
	assert.Len(users, 1)
}

func TestUserData_ReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	mock.ExpectQuery(readAllUsersQuery).WillReturnError(errors.New("something went wrong in func ReadAll"))
	users, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(users)
}

func TestUserData_Read(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	rows := sqlmock.
		NewRows([]string{"id_user", "login", "password", "gender", "email",
			"last_name", "first_name", "birthday", "city_id"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityId)
	mock.ExpectQuery(readUserQuery).WillReturnRows(rows)
	user, err := data.Read(testUser.IdUser)
	assert.NoError(err)
	assert.NotEmpty(user)
	assert.Equal(user, testUser)
}

func TestUserData_ReadErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	mock.ExpectQuery(readUserQuery).
		WillReturnError(errors.New("something went wrong in func Read"))
	user, err := data.Read(testUser.IdUser)
	assert.Error(err)
	assert.Empty(user)
}

func TestUserData_Add(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	mock.ExpectBegin()
	mock.ExpectExec(insertUserQuery).
		WithArgs(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	id, err := data.Add(testUser)
	assert.NoError(err)
	assert.NotEmpty(id)
	assert.Equal(id, testUser.IdUser)
}

func TestUserData_AddErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	mock.ExpectBegin()
	mock.ExpectExec(insertUserQuery).
		WithArgs(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityId).
		WillReturnError(errors.New("something went wrong in func Add"))
	mock.ExpectCommit()
	id, err := data.Add(testUser)
	assert.Error(err)
	assert.Equal(-1, id)
}

func TestUserData_Update(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	mock.ExpectBegin()
	mock.ExpectExec(insertUserQuery).
		WithArgs(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityId).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	_, err = data.Update(testUser)
	assert.NoError(err)
}

func TestUserData_UpdateErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	mock.ExpectQuery(readAllUsersWhereQuery).WithArgs(testUser.IdUser).
		WillReturnError(errors.New("something went wrong in func Update"))
	_, err = data.Update(testUser)
	assert.Error(err)
}

func TestUserData_Delete(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	rows := sqlmock.
		NewRows([]string{"id_user", "login", "password", "gender", "email",
			"last_name", "first_name", "birthday", "city_id"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityId)
	mock.ExpectQuery(readAllUsersWhereQuery).WithArgs(testUser.IdUser).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(deleteUserQuery).WithArgs(testUser.IdUser).WillReturnResult(sqlmock.NewResult(
		0, 1))
	mock.ExpectCommit()
	err = data.Delete(3)
	assert.NoError(err)
}

func TestUserData_DeleteErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserData(dbGorm)
	rows := sqlmock.
		NewRows([]string{"id_user", "login", "password", "gender", "email",
			"last_name", "first_name", "birthday", "city_id"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityId)
	mock.ExpectQuery(readAllUsersWhereQuery).WithArgs(testUser.IdUser).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(deleteUserQuery).WithArgs(testUser.IdUser).
		WillReturnError(errors.New("something went wrong in func Delete"))
	mock.ExpectCommit()
	err = data.Delete(3)
	assert.Error(err)
}
