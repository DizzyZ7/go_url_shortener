package postgres

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq" // Драйвер PostgreSQL
)

// DB представляет клиент для работы с базой данных PostgreSQL.
type DB struct {
    *sql.DB
}

// NewDB устанавливает соединение с базой данных и возвращает клиент.
func NewDB(dbURL string) *DB {
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }
    if err = db.Ping(); err != nil {
        log.Fatalf("Не удалось проверить соединение с базой данных: %v", err)
    }
    log.Println("Успешное подключение к PostgreSQL!")
    return &DB{db}
}

// Migrate создает необходимую таблицу, если она еще не существует.
func (d *DB) Migrate() {
    query := `
    CREATE TABLE IF NOT EXISTS urls (
        id SERIAL PRIMARY KEY,
        short_url VARCHAR(255) NOT NULL UNIQUE,
        original_url TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        clicks INT DEFAULT 0
    );`

    _, err := d.Exec(query)
    if err != nil {
        log.Fatalf("Не удалось выполнить миграцию: %v", err)
    }
    log.Println("Миграция PostgreSQL выполнена успешно.")
}
