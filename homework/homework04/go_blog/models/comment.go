package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	PostID    uint           `json:"post_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Post Post `json:"Post,omitempty" gorm:"foreignKey:PostID"`
}

type CreateCommentRequest struct {
	PostID  uint   `json:"postID"  binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint   `json:"userID"`
}

type UpdateCommentRequest struct {
	Content string `json:"content binding:"required"`
}

type CommentRequest struct {
	ID string `json:"id" binding:"required"`
}

type CommentPageRequest struct {
	page     uint `json:"page" binding:"required"`
	pageSize uint `json:"pageSize" binding:"required"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"userID"`
	PostID    uint      `json:"PostID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
