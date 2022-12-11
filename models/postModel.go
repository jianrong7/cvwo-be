package models

import (
	"github.com/jianrong/cvwo-be/initializers"
	"gorm.io/gorm"
)

type Post struct {
  gorm.Model
  Title string `gorm:"type:text" json:"title"`
	Content string `gorm:"type:text" json:"content"`
  // Tags pq.StringArray `gorm:type:text[]"`
  UserId uint
  User User `gorm:"foreignKey:UserId" json:"user"`
  Upvotes uint `gorm:"size:255" json:"upvotes"`
  Downvotes uint `gorm:"size:255" json:"downvotes"`
  Comments []Comment
}

func (post *Post) Save() (*Post, error) {
  err := initializers.DB.Create(&post).Error
  if err != nil {
    return &Post{}, err
  }
  return post, nil
}
// func FindUserByUsername(username string) (User, error) {
// 	var user User
// 	err := initializers.DB.Where("username=?", username).Find(&user).Error
// 	// initializers.DB.First(&user, "username = ?", body.Username)
// 	if err != nil {
// 		return User{}, err
// 	}
// 	return user, nil
// }
