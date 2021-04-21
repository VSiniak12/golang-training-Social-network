package data

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var testUserCity = &UserCity{
	IdUser:    3,
	Login:     "MockLogin",
	Password:  "MockPassword",
	Gender:    true,
	Email:     "mock@mail.ru",
	LastName:  "MockLastName",
	FirstName: "MockFirstName",
	Birthday:  time.Date(2000, 3, 10, 0, 0, 0, 0, time.UTC),
	CityName:  "Moskva",
}

func TestUserCityData_ReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserCityData(dbGorm)
	rows := sqlmock.
		NewRows([]string{"id_user", "login", "password", "gender", "email",
			"last_name", "first_name", "birthday", "name"}).
		AddRow(testUserCity.IdUser, testUserCity.Login, testUserCity.Password, testUserCity.Gender,
			testUserCity.Email, testUserCity.LastName, testUserCity.FirstName, testUserCity.Birthday,
			testUserCity.CityName)
	mock.ExpectQuery(readAllUserCityQuery).WillReturnRows(rows)
	users, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(users)
	assert.Equal(users[0], *testUserCity)
	assert.Len(users, 1)
}

func TestUserCityData_ReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	dbGorm, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println("Error with connection: ", err)
	}
	data := NewUserCityData(dbGorm)
	mock.ExpectQuery(readAllUserCityQuery).
		WillReturnError(errors.New("something went wrong in func ReadAll"))
	users, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(users)
}
