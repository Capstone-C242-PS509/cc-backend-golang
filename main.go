package main

import (
	"backend-bangkit/handler"
	"backend-bangkit/infra/postgresql"
	"backend-bangkit/pkg/gcloud"
	"backend-bangkit/repository"
	service "backend-bangkit/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := postgresql.GetDBInstance()
	gcsUploader, err := gcloud.NewGCSUploader()
	if err != nil {
		log.Fatalf(err.Message())
	}

	// db.AutoMigrate(&entity.Article{}, &entity.Tagging{}, &entity.User{}, &entity.PredictionResult{})

	articleRepo := repository.NewArticleRepository(db)
	tagRepo := repository.NewTaggingRepository(db)
	userRepo := repository.NewUserPg(db)
	musicRepo := repository.NewMusicRepository(db)
	musicService := service.NewMusicService(musicRepo)
	musicHandler := handler.NewMusicHandler(musicService)
	articleService := service.NewArticleService(articleRepo, tagRepo, gcsUploader)
	articleHandler := handler.NewArticleHandler(articleService)
	taggingService := service.NewTaggingService(tagRepo)
	taggingHandler := handler.NewTaggingHandler(taggingService)
	userService := service.NewAuthService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := service.NewPredictionService(predictionRepo)
	predictionHandler := handler.NewPredictionHandler(predictionService)

	route := gin.Default()

	UsersRoute := route.Group("/user")
	{

		UsersRoute.POST("/register", userHandler.Register)
		UsersRoute.POST("/login", userHandler.Login)
	}

	ArticleRoute := route.Group("/articles")
	{

		ArticleRoute.POST("", userService.Authentication(), userService.AdminAuthorization(), articleHandler.CreateArticle)
		ArticleRoute.GET("", userService.Authentication(), articleHandler.GetAllArticles)
		ArticleRoute.DELETE("/:id", userService.Authentication(), userService.AdminAuthorization(), articleHandler.DeleteArticle)
	}

	TagsRoute := route.Group("/tags")
	{
		TagsRoute.POST("", userService.Authentication(), userService.AdminAuthorization(), taggingHandler.CreateTag)
		TagsRoute.GET("", userService.Authentication(), taggingHandler.GetAllTags)
		TagsRoute.PUT("/:id", userService.Authentication(), userService.AdminAuthorization(), taggingHandler.UpdateTag)
		TagsRoute.DELETE("/:id", userService.Authentication(), userService.AdminAuthorization(), taggingHandler.DeleteTag)
	}

	predictRoute := route.Group("/predictions")
	{

		predictRoute.POST("", userService.Authentication(), predictionHandler.SavePrediction)

		predictRoute.GET("/summary", userService.Authentication(), predictionHandler.GetPredictionSummary)
	}

	musicRoute := route.Group("/music")
	{

		musicRoute.GET("", userService.Authentication(), musicHandler.GetRandomMusic)
	}

	route.Run(":8080")
}
