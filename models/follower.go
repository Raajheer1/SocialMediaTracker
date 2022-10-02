package models

import (
	"errors"
	"time"
)

type Follower struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Follower  uint `gorm:"size:255;not null" json:"follower"`
	AccountID uint
}

func GetFollowerByID(aID uint) (Follower, error) {
	var f Follower

	if err := DB.Find(&f, aID).Error; err != nil {
		return f, errors.New("Error fetching followers!")
	}

	return f, nil
}

func (f *Follower) SaveFollower() (*Follower, error) {
	var err error
	err = DB.Create(&f).Error
	if err != nil {
		return &Follower{}, err
	}

	return f, nil
}
