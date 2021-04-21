package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

type UserCityData struct {
	db *gorm.DB
}

type UserCity struct {
	IdUser    int `gorm:"primaryKey;column:id_user"`
	Login     string
	Password  string
	Gender    bool
	Email     string
	LastName  string
	FirstName string
	Birthday  time.Time
	CityName  string `gorm:"column:name"`
}

func NewUserCityData(db *gorm.DB) *UserCityData {
	return &UserCityData{db: db}
}

func (u UserCityData) ReadAll() ([]UserCity, error) {
	var users []UserCity
	result := u.db.Select("users.id_user, users.login, users.password, users.gender," +
		"users.email, users.last_name, users.first_name, users.birthday, cities.name").
		Joins(joinNameCity).Table("social_network.users").Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read users_city from database, error: %w", result.Error)
	}
	return users, nil
}
