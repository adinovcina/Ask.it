package main

import (
	"github.com/adinovcina/config"
	"github.com/adinovcina/controller"
	"github.com/adinovcina/middleware"
	"github.com/adinovcina/repository"
	"github.com/adinovcina/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	postRepository repository.PostRepository = repository.NewPostRepository(db)
	userRepository repository.UserRepository = repository.NewUserRepository(db)

	jwtService  service.JWTService  = service.NewJWTService()
	postService service.PostService = service.NewPostService(postRepository)
	userService service.UserService = service.NewUserService(userRepository)

	postController controller.PostController = controller.NewPostController(postService, jwtService)
	authController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	postRoutes := r.Group("api/post", middleware.AuthorizeJWT(jwtService))
	{
		postRoutes.GET("/", postController.GetAll)
		postRoutes.POST("/", postController.Insert)
	}

	r.Run()
}
