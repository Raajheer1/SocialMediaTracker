package models

import (
	"errors"
	"time"
)

type Account struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Handle     string `gorm:"size:255;not null;unique" json:"handle"`
	Platform   string `gorm:"size:255;not null;" json:"platform"`
	Department string `gorm:"size:255;not null;" json:"department"`
}

func GetAccountByID(aid uint) (Account, error) {
	var a Account

	if err := DB.First(&a, aid).Error; err != nil {
		return a, errors.New("Account not found!")
	}

	return a, nil

}

func GetAccountByHandle(aHandle string, aPlatform string) (Account, error) {
	var a Account

	if err := DB.Where("Handle = ? AND Platform >= ?", aHandle, aPlatform).First(&a).Error; err != nil {
		return a, errors.New("Account not found!")
	}

	return a, nil
}

func (a *Account) SaveAccount() (*Account, error) {
	var err error
	err = DB.Create(&a).Error
	if err != nil {
		return &Account{}, err
	}
	return a, nil
}
