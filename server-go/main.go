package main

import (
	"os"

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
	db                     *gorm.DB                          = config.SetupDatabaseConnection()
	postRepository         repository.PostRepository         = repository.NewPostRepository(db)
	userRepository         repository.UserRepository         = repository.NewUserRepository(db)
	answerRepository       repository.AnswerRepository       = repository.NewAnswerRepository(db)
	userPostRepository     repository.UserPostRepository     = repository.NewUserPostRepository(db)
	answerPostRepository   repository.AnswerPostRepository   = repository.NewAnswerPostRepository(db)
	notificationRepository repository.NotificationRepository = repository.NewNotificationRepository(db)

	jwtService          service.JWTService          = service.NewJWTService()
	postService         service.PostService         = service.NewPostService(postRepository)
	userService         service.UserService         = service.NewUserService(userRepository)
	answerService       service.AnswerService       = service.NewAnswerService(answerRepository)
	userPostService     service.UserPostService     = service.NewUserPostService(userPostRepository)
	answerPostService   service.AnswerPostService   = service.NewAnswerPostService(answerPostRepository)
	notificationService service.NotificationService = service.NewNotificationService(notificationRepository)

	postController         controller.PostController         = controller.NewPostController(postService, jwtService, userPostService)
	authController         controller.UserController         = controller.NewUserController(userService, jwtService)
	answerController       controller.AnswerController       = controller.NewAnswerController(answerService, jwtService, answerPostService)
	postUserController     controller.UserPostController     = controller.NewUserPostController(userPostService, jwtService, postService)
	answerPostController   controller.AnswerPostController   = controller.AnswerUserPostController(answerPostService, jwtService, answerService)
	notificationController controller.NotificationController = controller.NewNotificationController(notificationService, jwtService)
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
		postProtectedRoutes.GET("/myPosts", postController.MyPosts)
	}

	postRoutes := r.Group("api/post")
	{
		postRoutes.GET("/", postController.GetAll)
		postRoutes.GET("/mostLikes", postController.MostLikedPost)
	}

	answerRoutes := r.Group("api/answer")
	{
		answerRoutes.GET("/", answerController.GetAll)
		answerRoutes.GET("/mostAnswers", answerController.MostAnswers)
	}

	answerProtectedRoutes := r.Group("api/answer", middleware.AuthorizeJWT(jwtService))
	{
		answerProtectedRoutes.POST("/", answerController.Insert)
		answerProtectedRoutes.PUT("/", answerController.EditAnswer)
		answerProtectedRoutes.DELETE("/:id", answerController.DeleteAnswer)
	}

	gradePostRoutes := r.Group("api/grade")
	{
		gradePostRoutes.GET("/", postUserController.GetAll)
	}

	gradePostProtectedRoutes := r.Group("api/grade", middleware.AuthorizeJWT(jwtService))
	{
		gradePostProtectedRoutes.POST("/", postUserController.Insert)
		gradePostProtectedRoutes.PUT("/", postUserController.Update)
	}

	gradeAnswerRoutes := r.Group("api/answer/grade")
	{
		gradeAnswerRoutes.GET("/", answerPostController.GetAll)
	}

	gradeAnswerProtectedRoutes := r.Group("api/answer/grade", middleware.AuthorizeJWT(jwtService))
	{
		gradeAnswerProtectedRoutes.POST("/", answerPostController.Insert)
		gradeAnswerProtectedRoutes.PUT("/", answerPostController.Update)
	}

	passwordRoutes := r.Group("api/changepassword", middleware.AuthorizeJWT(jwtService))
	{
		passwordRoutes.POST("/", authController.ChangePassword)
	}

	notificationRoutes := r.Group("api/notification", middleware.AuthorizeJWT(jwtService))
	{
		notificationRoutes.GET("/", notificationController.GetAll)
	}

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
