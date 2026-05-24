package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_blog/models"
	"go_blog/services"
	"go_blog/utils"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	commentService *services.CommentService
	userService    *services.UserService
	postService    *services.PostService
	jwtSecret      []byte
}

func NewCommentHandler(userService *services.UserService, postService *services.PostService, commentService *services.CommentService, jwtSecret []byte) *CommentHandler {
	return &CommentHandler{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
		jwtSecret:      jwtSecret,
	}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var req models.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, parseValidationErrors(err))
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	user, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	if user == nil {
		utils.HandleError(c, errors.New("user not found"))
		return
	}
	req.UserID = userID.(uint)

	post, err := h.postService.GetPostByID(req.PostID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	if post == nil {
		utils.HandleError(c, errors.New("post not found"))
		return
	}

	comment, err := h.commentService.CreateComment(req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, models.CommentResponse{
		ID:      comment.ID,
		Content: comment.Content,
		UserID:  comment.UserID,
	})
}

// GetCommens 获取文章列表
func (h *CommentHandler) GetComments(c *gin.Context) {
	var req models.CommentPageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, parseValidationErrors(err))
		return
	}

	var comments []models.Comment

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	comments, err := h.commentService.GetComments(pageSize, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 获取总数
	var total int64
	total, err = h.commentService.GetGetCommentsTotal()

	utils.Success(c, gin.H{
		"comments":  comments,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}
