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
	db                   *gorm.DB                        = config.SetupDatabaseConnection()
	postRepository       repository.PostRepository       = repository.NewPostRepository(db)
	userRepository       repository.UserRepository       = repository.NewUserRepository(db)
	answerRepository     repository.AnswerRepository     = repository.NewAnswerRepository(db)
	userPostRepository   repository.UserPostRepository   = repository.NewUserPostRepository(db)
	answerPostRepository repository.AnswerPostRepository = repository.NewAnswerPostRepository(db)

	jwtService        service.JWTService        = service.NewJWTService()
	postService       service.PostService       = service.NewPostService(postRepository)
	userService       service.UserService       = service.NewUserService(userRepository)
	answerService     service.AnswerService     = service.NewAnswerService(answerRepository)
	userPostService   service.UserPostService   = service.NewUserPostService(userPostRepository)
	answerPostService service.AnswerPostService = service.NewAnswerPostService(answerPostRepository)

	postController       controller.PostController       = controller.NewPostController(postService, jwtService)
	authController       controller.UserController       = controller.NewUserController(userService, jwtService)
	answerController     controller.AnswerController     = controller.NewAnswerController(answerService, jwtService)
	postUserController   controller.UserPostController   = controller.NewUserPostController(userPostService, jwtService, postService)
	answerPostController controller.AnswerPostController = controller.AnswerUserPostController(answerPostService, jwtService, answerService)
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

	postProtectedRoutes := r.Group("api/post", middleware.AuthorizeJWT(jwtService))
	{
		postProtectedRoutes.POST("/", postController.Insert)
	}

	postRoutes := r.Group("api/post")
	{
		postRoutes.GET("/", postController.GetAll)
	}

	answerRoutes := r.Group("api/answer")
	{
		answerRoutes.GET("/", answerController.GetAll)
		// answerRoutes.POST("/", answerController.Insert)
		// answerRoutes.GET("/mostAnswers", answerController.MostAnswers)
		// answerRoutes.PUT("/", answerController.UpdateAnswerMark)
	}

	answerProtectedRoutes := r.Group("api/answer", middleware.AuthorizeJWT(jwtService))
	{
		answerProtectedRoutes.POST("/", answerController.Insert)
		// answerRoutes.GET("/mostAnswers", answerController.MostAnswers)
		// answerRoutes.PUT("/", answerController.UpdateAnswerMark)
	}

	gradePostRoutes := r.Group("api/grade")
	{
		gradePostRoutes.GET("/", postUserController.GetAll)
		// gradeRoutes.POST("/", postUserController.Insert)
		// gradeRoutes.PUT("/", postUserController.Put)
	}

	gradePostProtectedRoutes := r.Group("api/grade", middleware.AuthorizeJWT(jwtService))
	{
		gradePostProtectedRoutes.POST("/", postUserController.Insert)
		// gradeRoutes.PUT("/", postUserController.Put)
	}

	gradeAnswerRoutes := r.Group("api/answer/grade")
	{
		gradeAnswerRoutes.GET("/", answerPostController.GetAll)
	}

	gradeAnswerProtectedRoutes := r.Group("api/answer/grade", middleware.AuthorizeJWT(jwtService))
	{
		gradeAnswerProtectedRoutes.POST("/", answerPostController.Insert)
	}

	r.Run()
}
