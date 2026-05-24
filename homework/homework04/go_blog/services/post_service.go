package services

import (
	"errors"
	"go_blog/models"
	"go_blog/utils"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) CreatePost(req models.CreatePostRequest) (*models.Post, error) {
	// 创建博客
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := s.db.Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppError(404, "Post not found")
		}
		return nil, err
	}
	return &post, nil
}

// 1. 函数命名更语义化：根据用户ID获取分页文章列表
// 2. 接收分页参数，而不是硬编码 pageSize/offset
// 3. 返回 *[]Post 是错误的，直接返回 []Post 更标准
func (s *PostService) GetPostsByUserID(userID uint, pageSize int, offset int) ([]models.Post, error) {
	var posts []models.Post

	// 构建查询：只查当前用户的文章 + 预加载用户信息
	query := s.db.Model(&models.Post{}).
		Where("user_id = ?", userID). // 关键：你漏了用户ID筛选！
		Preload("User").              // 预加载关联用户
		Order("created_at DESC")      // 按创建时间倒序

	// 分页（如果有传分页参数才加）
	if pageSize > 0 {
		query = query.Limit(pageSize)
	}
	if offset >= 0 {
		query = query.Offset(offset)
	}

	// 执行查询
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPostsTotalByUserID 根据用户ID获取文章总数
func (s *PostService) GetPostsTotalByUserID(userID uint) (int64, error) {
	var total int64
	err := s.db.
		Model(&models.Post{}).
		Where("user_id = ?", userID).
		Count(&total).Error

	return total, err
}

// UpdatePost 更新文章
func (s *PostService) GetUserByIDAndUserID(id uint64, userID uint) (*models.Post, error) {
	var post models.Post

	// ✅ 正确：同时按 文章ID + 所属用户ID 查询
	err := s.db.Where("id = ? AND user_id = ?", id, userID).First(&post).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppError(404, "Post not found")
		}
		return nil, err
	}

	return &post, nil
}

// UpdatePost 更新文章
func (s *PostService) UpdatePost(postID uint, content string) error {
	return s.db.Model(&models.Post{}).
		Where("id = ?", postID).
		Update("content", content).
		Error
}

// DeletePost 删除文章
func (s *PostService) DeletePost(id uint, userID uint) error {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		return err
	}

	// 检查是否是文章作者
	if post.UserID != userID {
		return utils.NewAppError(409, "this post is not user's post")
	}

	// 删除文章（软删除）
	if err := s.db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
