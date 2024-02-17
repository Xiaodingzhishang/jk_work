package main

import (
	"github.com/Xiaodingzhishang/jk_work/webook/internal/repository"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/repository/dao"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/service"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/web"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDb()
	server := initWebServer()
	u := initUser(db)
	u.RegisterRoutes(server)
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			//return origin == "https://github.com"
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "www.xiaoding.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	//store := cookie.NewStore([]byte("secret"))
	//store := memstore.NewStore([]byte("Mwz4s3qHZjMpXJdUTn7QT5TZ7ztWWWaY"), []byte("2mXLah2yEOynD2BESJwe0xHuz3TDG8LG"))
	store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("Mwz4s3qHZjMpXJdUTn7QT5TZ7ztWWWaY"), []byte("2mXLah2yEOynD2BESJwe0xHuz3TDG8LG"))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("mysession", store))
	server.Use(middleware.NewLoginJwtMiddlewareBuilder().IgnorePaths("/users/login").IgnorePaths("/users/signup").Builder())
	return server
}

func initUser(db *gorm.DB) *web.UserHandle {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandle(svc)
	return u
}
func initDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/webook?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		//panic 相当于整个goroutine结束
		//一旦初始化过程出错，应用就不要启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
