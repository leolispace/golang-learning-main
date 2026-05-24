package services

import (
	"go_blog/models"
	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

func (s *CommentService) CreateComment(req models.CreateCommentRequest) (*models.Comment, error) {
	// 创建博客
	comment := models.Comment{
		Content: req.Content,
		UserID:  req.UserID,
		PostID:  req.PostID,
	}

	if err := s.db.Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

// 1. 函数命名更语义化：根据用户ID获取分页文章列表
// 2. 接收分页参数，而不是硬编码 pageSize/offset
// 3. 返回 *[]Post 是错误的，直接返回 []Post 更标准
func (s *CommentService) GetComments(pageSize int, offset int) ([]models.Comment, error) {
	var comments []models.Comment

	// 构建查询：只查当前用户的文章 + 预加载用户信息
	query := s.db.Model(&models.Comment{}).
		Preload("User").         // 预加载关联用户
		Order("created_at DESC") // 按创建时间倒序// 构建查询：只查当前用户的文章 + 预加载用户信息

	// 分页（如果有传分页参数才加）
	if pageSize > 0 {
		query = query.Limit(pageSize)
	}
	if offset >= 0 {
		query = query.Offset(offset)
	}

	// 执行查询
	if err := query.Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

// GetPostsTotalByUserID 根据用户ID获取文章总数
func (s *CommentService) GetGetCommentsTotal() (int64, error) {
	var total int64
	err := s.db.
		Model(&models.Comment{}).
		Count(&total).Error

	return total, err
}
