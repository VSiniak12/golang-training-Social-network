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

var testUser = User{
	IdUser:    3,
	Login:     "MockLogin",
	Password:  "MockPassword",
	Gender:    true,
	Email:     "mock@mail.ru",
	LastName:  "MockLastName",
	FirstName: "MockFirstName",
	Birthday:  time.Date(2000, 3, 10, 0, 0, 0, 0, time.UTC),
	CityName:  "MockCity",
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
			"last_name", "first_name", "birthday", "name"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityName)
	mock.ExpectQuery(readAllUsersQuery).WillReturnRows(rows)
	users, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(users)
	assert.Equal(users[0], testUser)
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
			"last_name", "first_name", "birthday", "name"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityName)
	mock.ExpectQuery(readUsersQuery).WillReturnRows(rows)
	user, err := data.Read(testUser.IdUser)
	assert.NoError(err)
	assert.NotEmpty(user)
	assert.Equal(*user, testUser)
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
	rows := sqlmock.
		NewRows([]string{"id_user", "login", "password", "gender", "email",
			"last_name", "first_name", "birthday", "name"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityName)
	mock.ExpectQuery(readAllUsersUpdateQuery).WithArgs(testUser.IdUser).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(updateUsersQuery).WithArgs(testUser.Login, testUser.IdUser).WillReturnResult(sqlmock.NewResult(
		0, 1))
	mock.ExpectCommit()
	err = data.Update(3, "Slava")
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
	mock.ExpectQuery(readAllUsersUpdateQuery).WithArgs(testUser.IdUser).
		WillReturnError(errors.New("something went wrong in func Update..."))
	err = data.Update(3, "Slava")
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
			"last_name", "first_name", "birthday", "name"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityName)
	mock.ExpectQuery(readAllUsersUpdateQuery).WithArgs(testUser.IdUser).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(deleteUsersQuery).WithArgs(testUser.IdUser).WillReturnResult(sqlmock.NewResult(
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
			"last_name", "first_name", "birthday", "name"}).
		AddRow(testUser.IdUser, testUser.Login, testUser.Password, testUser.Gender, testUser.Email,
			testUser.LastName, testUser.FirstName, testUser.Birthday, testUser.CityName)
	mock.ExpectQuery(readAllUsersUpdateQuery).WithArgs(testUser.IdUser).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(deleteUsersQuery).WithArgs(testUser.IdUser).
		WillReturnError(errors.New("something went wrong in func Delete"))
	mock.ExpectCommit()
	err = data.Delete(3)
	assert.Error(err)
}
