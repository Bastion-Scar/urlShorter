package main

import (
	"awesomeProject13/internal/logger"
	"awesomeProject13/internal/middlewares"
	"awesomeProject13/internal/service"
	"awesomeProject13/internal/storage"
	"awesomeProject13/models"
	"github.com/gin-gonic/gin"
)

//TODO: Написать тесты, чтоб покрытие было 80% и более
//TODO: Написать CI файл для GitHub Actions
//TODO: ДОПОЛНИТЕЛЬНО написать Dockerfile и Docker-compose

func main() {
	zapLogger := logger.Logger()
	initStorage := storage.NewMySQLStorage()
	userService := service.NewUserService(initStorage)

	r := gin.New()
	db := storage.InitDB()

	r.Use(middlewares.LoggerMiddleware(zapLogger))
	r.Use(gin.Recovery())
	r.POST("/shorten", func(c *gin.Context) {
		userService.Save(c)
	})

	r.GET("/:code", func(c *gin.Context) {
		userService.Load(c)

	})

	go models.SendLogs(db)

	if err := r.Run(":8080"); err != nil {
		zapLogger.Fatal("Ошибка при запуске сервера")
	}

}
