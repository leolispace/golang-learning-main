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

type PostHandler struct {
	postService *services.PostService
	userService *services.UserService
	jwtSecret   []byte
}

func NewPostHandler(userService *services.UserService, postService *services.PostService, jwtSecret []byte) *PostHandler {
	return &PostHandler{
		userService: userService,
		postService: postService,
		jwtSecret:   jwtSecret,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req models.CreatePostRequest
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

	post, err := h.postService.CreatePost(req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, models.PostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	})
}

func (h *PostHandler) GetPostInfo(c *gin.Context) {
	postIDStr := c.Query("postID")
	// 2. 判断是否为空
	if postIDStr == "" {
		c.JSON(400, gin.H{"error": "postID 不能为空"})
		return
	}

	// 3. 字符串转 uint（Go 标准库）
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "postID 必须是数字"})
		return
	}

	// 4. 转成 uint（如果你需要的是 uint 而不是 uint64）
	uid := uint(postID)

	post, err := h.postService.GetPostByID(uid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, models.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	})
}

// GetPosts 获取文章列表
func (h *PostHandler) GetPosts(c *gin.Context) {
	var req models.PostPageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, parseValidationErrors(err))
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var posts []models.Post

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

	posts, err := h.postService.GetPostsByUserID(userID.(uint), pageSize, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 获取总数
	var total int64
	total, err = h.postService.GetPostsTotalByUserID(userID.(uint))

	utils.Success(c, gin.H{
		"posts":     posts,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}
func (h *PostHandler) UpdatePost(c *gin.Context) {
	// 1. 获取评论ID
	idStr := c.Query("id")
	postID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "ID格式错误")
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	user, err := h.postService.GetUserByIDAndUserID(postID, userID.(uint))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	if user == nil {
		utils.HandleError(c, errors.New("comment not found by commentid and userid"))
		return
	}

	// 2. 获取前端传的新内容
	var req struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 3. 执行更新
	err = h.postService.UpdatePost(uint(postID), req.Content)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新失败")
		return
	}
	utils.Success(c, gin.H{"msg": "更新成功"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "ID格式错误")
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 3. 执行更新
	err = h.postService.DeletePost(uint(postID), userID.(uint))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新失败")
		return
	}
	utils.Success(c, gin.H{"msg": "更新成功"})
}
