package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	UserID    uint           `json:"userID" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=20"`
	Content string `json:"content" binding:"required"`
	UserID  uint   `json:"userID"`
}

type UpdatePostRequest struct {
	ID      string `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required,min=3,max=20"`
	Content string `json:"content" binding:"required"`
}

type PostRequest struct {
	ID string `json:"id" binding:"required"`
}

type PostPageRequest struct {
	page     uint `json:"page" binding:"required"`
	pageSize uint `json:"pageSize" binding:"required"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"userID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
