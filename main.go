package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Endropr/ai-programming-mentor/internal/adapter/repository"
	"github.com/Endropr/ai-programming-mentor/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Загружаем переменные из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// 2. Берем строку подключения
	dbURL := os.Getenv("DB_URL")

	// 3. Подключаемся к базе
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе: %v", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("✅ Успешно подключились к PostgreSQL!")

	// 4. Проверяем работу нашего репозитория (Адаптера)
	repo := repository.NewPostgresRepo(conn)
	
	testMsg := domain.Message{
		UserID:  777,
		Role:    "user",
		Content: "Привет из кода на Go!",
	}

	err = repo.SaveMessage(context.Background(), testMsg)
	if err != nil {
		log.Fatalf("Не удалось сохранить сообщение: %v", err)
	}

	fmt.Println("🚀 Тестовое сообщение успешно сохранено в базе!")
}


