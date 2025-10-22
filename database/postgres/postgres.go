package postgres

import (
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // Импорт драйвера PostgreSQL
)

// DB содержит подключение к базе данных.
var DB *sqlx.DB

// InitDB устанавливает соединение с базой данных.
func InitDB(connStr string) {
    var err error
    DB, err = sqlx.Connect("postgres", connStr)
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }

    // Проверяем соединение
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Не удалось проверить соединение с базой данных: %v", err)
    }

    log.Println("Успешно подключились к базе данных PostgreSQL!")
}

// MigrateDB выполняет миграции для создания необходимых таблиц.
func MigrateDB() {
    query := `
    CREATE TABLE IF NOT EXISTS urls (
        id SERIAL PRIMARY KEY,
        short_url TEXT UNIQUE NOT NULL,
        original_url TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        clicks INT DEFAULT 0
    );
    `
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatalf("Не удалось выполнить миграцию: %v", err)
    }
    log.Println("Миграция базы данных успешно выполнена.")
}
