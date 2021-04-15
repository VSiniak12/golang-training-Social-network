package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

type User struct {
	IdUser    int `gorm:"primaryKey;column:id_user"`
	Login     string
	Password  string
	Gender    bool
	Email     string
	LastName  string
	FirstName string
	Birthday  time.Time
	CityId    int
	CityName  string `gorm:"column:name"`
}

func NewUserData(db *gorm.DB) *UserData {
	return &UserData{db: db}
}

func (u UserData) ReadAll() ([]User, error) {
	var users []User
	result := u.db.Select("users.id_user, users.login, users.password, users.gender," +
		"users.email, users.last_name, users.first_name, users.birthday, cities.name").
		Joins(joinNameCity).Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read users from database, error: %w", result.Error)
	}
	return users, nil
}

func (u UserData) Read(id int) (*User, error) {
	var user User
	result := u.db.Select("users.id_user, users.login, users.password, users.gender,"+
		"users.email, users.last_name, users.first_name, users.birthday, cities.name").
		Joins(joinNameCity).Where("id_user = ?", id).Find(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read user from database, error: %w", result.Error)
	}
	if user.IdUser == 0 {
		return nil, fmt.Errorf("don't have user with id equals: %d", id)
	}
	return &user, nil
}

func (u UserData) Delete(id int) error {
	var user User
	result := u.db.Where("id_user = ?", id).Find(&user).Delete(&user)
	if result.Error != nil {
		return fmt.Errorf("can't delete user to database, error: %w", result.Error)
	}
	if user.IdUser == 0 {
		return fmt.Errorf("don't have user with id equals: %d", id)
	}
	return nil
}

func (u UserData) Update(id int, value string) error {
	var user User
	result := u.db.Where("id_user = ?", id).Find(&user)
	if result.Error != nil {
		return fmt.Errorf("can't update user to database, error: %w", result.Error)
	}
	if user.IdUser == 0 {
		return fmt.Errorf("don't have user with id equals: %d", id)
	}
	user.Login = value
	u.db.Omit("name").Save(&user)
	return nil
}

func (u UserData) Add(user User) (int, error) {
	u.db.Table("social_network.cities").Where("name = ?", user.CityName).Select("id_city").Find(&user.CityId)
	if user.CityId == 0 {
		return -1, fmt.Errorf("city with name %s has not", user.CityName)
	}
	result := u.db.Omit("name").Create(&user)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create user to database, error: %w", result.Error)
	}
	return user.IdUser, nil
}

func (user User) String() string {
	return fmt.Sprintf("%d ,%s, %s, %v, %s, %s, %s, %s, %s; \n",
		user.IdUser, user.Login, user.Password, user.Gender, user.Email, user.LastName,
		user.FirstName, parseDate(user.Birthday), user.CityName)
}

func parseDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
