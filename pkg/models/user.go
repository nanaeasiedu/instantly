package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/ngenerio/instantly/pkg/utils"
	"github.com/ngenerio/instantly/pkg/web/payloads"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int
	EmailAddress    string
	Password        string
	PasswordHash    string
	Token           string
	CallbackURL     string
	CurrentBalance  float64
	MobileNumber    string
	NetworkOperator string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Transactions    []Transaction `gorm:"ForeignKey:UserID"`
}

func CreateUser(user *payloads.User) (*User, error) {
	newUser := new(User)
	newUser.EmailAddress = user.Email
	newUser.Password = user.Password

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	newUser.PasswordHash = string(hashPassword)
	newUser.Token = utils.GenerateSecureKey()
	newUser.CurrentBalance = 10.00
	newUser.CreatedAt = time.Now()

	if err := db.Create(newUser).Error; err != nil {
		return nil, err
	}

	return newUser, err
}

func (u *User) GetUser(query map[string]interface{}) error {
	err := db.Where(query).First(u).Error
	return err
}

func DoesUserExist(query map[string]interface{}) (bool, error) {
	u := User{}
	err := db.Where(query).First(&u).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err == nil {
		return true, nil
	}
	return false, err
}
