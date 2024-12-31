package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wrlin1218/url_shortener/config"
	"github.com/wrlin1218/url_shortener/internal/controller"
	redis "github.com/wrlin1218/url_shortener/internal/dal/kv/impl"
	"github.com/wrlin1218/url_shortener/internal/dal/rdb"
	repo_impl "github.com/wrlin1218/url_shortener/internal/repo/impl"
	service_impl "github.com/wrlin1218/url_shortener/internal/service/impl"
	zap_logger "github.com/wrlin1218/url_shortener/pkg/logger/zap"
	"log"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("../config", "config.yaml")
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	log.Println("Configuration loaded successfully")

	// 初始化 basic tools
	zap_logger.Init(cfg.Log)
	db := rdb.Init(cfg.Database)
	cient := redis.NewCient(cfg.Redis)

	// 初始化 repo * service * controller
	userRepoImpl := repo_impl.NewUserRepoImpl(db, cient)
	linkRepoImpl := repo_impl.NewLinkRepoImpl(db, cient)
	userServiceImpl := service_impl.NewUserService(userRepoImpl, linkRepoImpl)
	linkServiceImpl := service_impl.NewLinkService(linkRepoImpl, userRepoImpl)

	r := gin.Default()
	// 初始化controller
	userController := controller.UserController{UserService: userServiceImpl}
	linkController := controller.LinkController{LinkService: linkServiceImpl}

	// 设置路由
	// 静态页面
	r.GET("/create", func(c *gin.Context) {
		c.File("../static/index.html") // 根据你的文件结构调整路径
	})
	// 后端请求
	r.GET("/user/create", userController.CreateUser)
	r.GET("/user/links", userController.GetAllLinksByUserName)
	r.POST("/link/create", linkController.CreateShortLink)
	r.GET("/link/delete", linkController.DeleteShortLink)
	r.GET("/:short_code", linkController.RedirectToOriginal)

	// 启动服务器
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
