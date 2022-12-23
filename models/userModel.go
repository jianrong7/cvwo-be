package models

import (
	"html"
	"strings"

	"github.com/jianrong/cvwo-be/initializers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
  gorm.Model
  Username string 	 `gorm:"size:255;not null;unique" json:"username"`
	Password string 	 `gorm:"size:255;not null;" json:"-"`
	Posts 	 []Post 	 `json:"posts"`
	Comments []Comment `json:"comments"`
	Ratings	 []Rating  `json:"ratings"`
}

func (user *User) Save() (*User, error) {
	err := initializers.DB.Omit(clause.Associations).Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// (pre-defined GORM hook) trim whitespace in username and hash password
func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := initializers.DB.Where("username=?", username).Find(&user).Error
	// initializers.DB.First(&user, "username = ?", body.Username)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(userId uint) (User, error) {
	var user User
	err := initializers.DB.Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at DESC")
	}).
	Preload("Ratings", func(db *gorm.DB) *gorm.DB {
		return db.Order("ratings.created_at DESC")
	}).
	Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Order("posts.created_at DESC")
	}).
	Where("ID=?", userId).
	Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}