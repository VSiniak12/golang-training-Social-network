package data

import (
	"fmt"

	"gorm.io/gorm"
)

type StateData struct {
	db *gorm.DB
}

type State struct {
	IdState   int `gorm:"primaryKey;column:id_state"`
	Name      string
	CountryId int
}

func NewStateData(db *gorm.DB) *StateData {
	return &StateData{db: db}
}
func (s StateData) ReadAll() ([]State, error) {
	var states []State
	result := s.db.Find(&states)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read states from database, error: %w", result.Error)
	}
	return states, nil
}

func (s StateData) Add(state State) (int, error) {
	result := s.db.Create(&state)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create state to database, error: %w", result.Error)
	}
	return state.IdState, nil
}

func (s StateData) Delete(id int) error {
	var state State
	result := s.db.Where("id_state = ?", id).Find(&state).Delete(&state)
	if result.Error != nil {
		return fmt.Errorf("can't delete state to database, error: %w", result.Error)
	}
	if state.IdState == 0 {
		return fmt.Errorf("don't have state with id equals: %d", id)
	}
	return nil
}

func (s StateData) Update(id int, name string) error {
	var state State
	result := s.db.Where("id_state = ?", id).Find(&state)
	if result.Error != nil {
		return fmt.Errorf("can't delete state to database, error: %w", result.Error)
	}
	if state.IdState == 0 {
		return fmt.Errorf("don't have state with id equals: %d", id)
	}
	state.Name = name
	s.db.Save(&state)
	return nil
}
