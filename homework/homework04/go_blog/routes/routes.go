package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/config"
	"go_blog/handlers"
	"go_blog/middleware"
	"go_blog/models"
	"go_blog/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// SetupRoutes 设置路由
func SetupRoutes() *gin.Engine {
	// 加载配置
	cfg := config.Load()

	// 构建MySQL连接字符串
	// ✅ 正确格式
	//dsn := cfg.Database.Username + ":" + cfg.Database.Password + "@tcp(127.0.0.1:3306)/" + cfg.Database.DBName + "?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	// 连接MySQL数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 初始化数据库
	//db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	//if err != nil {
	//	log.Fatalf("Failed to connect database: %v", err)
	//}

	// 自动迁移
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	//
	// 初始化服务
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService, []byte(cfg.JWT.Secret))

	postService := services.NewPostService(db)
	postHandler := handlers.NewPostHandler(userService, postService, []byte(cfg.JWT.Secret))

	commentService := services.NewCommentService(db)
	commentHandler := handlers.NewCommentHandler(userService, postService, commentService, []byte(cfg.JWT.Secret))

	r := gin.New()

	// 使用中间件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandlerMiddleware())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	// API路由组
	api := r.Group("/api/v1")
	{
		// 认证相关路由（无需认证）
		auth := api.Group("")
		{
			auth.POST("/users/login", userHandler.Login)
		}

		// 需要认证的路由
		authenticated := api.Group("")
		authenticated.Use(middleware.Auth([]byte(cfg.JWT.Secret)))
		{
			// 用户信息
			// 文章相关路由
			users := authenticated.Group("/users")
			{
				users.GET("/me", userHandler.GetProfile)
				users.PUT("/me", userHandler.UpdateProfile)
			}

			// 文章相关路由
			posts := authenticated.Group("/post")
			{
				posts.POST("/create", postHandler.CreatePost)
				posts.GET("/GetPostInfo", postHandler.GetPostInfo)
				posts.POST("/GetPosts", postHandler.GetPosts)
				posts.POST("/UpdatePost", postHandler.UpdatePost)
				posts.GET("/DeletePost", postHandler.DeletePost)

			}

			// 评论相关路由
			comments := authenticated.Group("/comments")
			{
				comments.POST("/CreateComment", commentHandler.CreateComment)
				comments.POST("/GetComments", commentHandler.GetComments)
			}
		}

		// 公开路由（无需认证）
		public := r.Group("/api/v1")
		{
			public.POST("/users/register", userHandler.Register)
		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Blog API is running",
		})
	})

	// 启动服务器
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return r
}
