# URL Shortener

Сервис для сокращения ссылок, написан на Go

## Возможности

- Сокращение URL и редирект
- Хранение ссылок в MySQL или в памяти
- Асинхронное логирование запросов
- REST API
- Unit-тесты
- CI через GitHub Actions

## Запуск

go run main.go
API
POST /shorten
Принимает JSON с URL, возвращает короткий код. (6 символов)

Пример:

{
  "url": "https://google.com"
}
Ответ:

{
  "shortURL": "ab1J45"
}
GET /:code
Перенаправляет на оригинальный URL по короткому коду.

Тесты

go test ./...

Переменные окружения
Создайте .env:

DSN=username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
Требования
Go 1.22
MySQL
github.com/gin-gonic/gin v1.10.1
github.com/joho/godotenv v1.5.1
github.com/stretchr/testify v1.10.0
go.uber.org/zap v1.27.0
gorm.io/driver/mysql v1.6.0
gorm.io/gorm v1.30.1