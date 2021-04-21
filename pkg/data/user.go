package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
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
}

func NewUserData(db *gorm.DB) *UserData {
	return &UserData{db: db}
}

func (u UserData) ReadAll() ([]User, error) {
	var users []User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read users from database, error: %w", result.Error)
	}
	return users, nil
}

func (u UserData) Read(id int) (*User, error) {
	var user User
	result := u.db.Where("id_user = ?", id).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read user from database, error: %w", result.Error)
	}
	return &user, nil
}

func (u UserData) Delete(id int) error {
	var user User
	result := u.db.Where("id_user = ?", id).Find(&user).Delete(&user)
	if result.Error != nil {
		return fmt.Errorf("can't delete user to database, error: %w", result.Error)
	}
	return nil
}

func (u UserData) Update(user *User) (*User, error) {
	result := u.db.Save(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("can't update user in database, error: %w", result.Error)
	}
	return user, nil
}

func (u UserData) Add(user *User) (int, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create user to database, error: %w", result.Error)
	}
	return user.IdUser, nil
}
