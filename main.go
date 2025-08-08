package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

//TODO: Распределить все по папкам
//TODO: Добавить Zap логгер и Middleware на его основе. Также запись в файл с lumberjack ротацией и настроить семплирование
//TODO: Заменить карту на MySQL, в таблице столбцы о запросе, IP, путь и т.п
//TODO: Добавить в .env все, что нужно
//TODO: Написать тесты, чтоб покрытие было 80% и более
//TODO: Написать CI файл для GitHub Actions
//TODO: ДОПОЛНИТЕЛЬНО написать Dockerfile и Docker-compose

type Storage interface {
	Save(code string, url string)
	Load(code string) (string, bool)
}

type MapStorage struct {
	data map[string]string
}

func NewMapStorage() *MapStorage {
	return &MapStorage{make(map[string]string)}
}

func (m *MapStorage) Save(code string, url string) {
	m.data[code] = url
}

func (m *MapStorage) Load(code string) (string, bool) {
	url, ok := m.data[code]
	return url, ok
}

type UserService struct {
	repo Storage
}

func NewUserService(repo Storage) *UserService {
	return &UserService{repo}
}

func (s *UserService) Save(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "binding failed",
		})
		return
	}

	shortenURL, err := getRandomURL()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot create shorten url",
		})
		return
	}

	s.repo.Save(shortenURL, req.URL)

	c.JSON(200, gin.H{
		"shortURL": shortenURL,
	})
}

func (s *UserService) Load(c *gin.Context) {
	code := c.Param("code")
	originalURL, ok := s.repo.Load(code)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "shorten code not found",
		})
		return
	}

	c.Redirect(302, originalURL)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomURL() (string, error) {
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code), nil
}

func main() {
	storage := NewMapStorage()
	userService := NewUserService(storage)

	r := gin.Default()
	r.POST("/shorten", func(c *gin.Context) {
		userService.Save(c)
	})

	r.GET("/:code", func(c *gin.Context) {
		userService.Load(c)

	})

	if err := r.Run(":8080"); err != nil {
		panic("Ошибка при запуске сервера")
	}
}
