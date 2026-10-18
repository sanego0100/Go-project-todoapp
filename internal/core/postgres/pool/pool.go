package core_postgres_pool

import (
	"context"
	"time"
)

// Pool — интерфейс пула соединений с базой данных.
// Конкретная реализация — в пакете pgx (internal/core/repository/postgres/pool/pgx).
//
// OpTimeout() возвращает максимальное время выполнения одного запроса к БД.
type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, error)
	Close()

	OpTimeout() time.Duration
}

// Rows — интерфейс для итерации по результатам SELECT-запроса (несколько строк).
type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

// Row — интерфейс для получения одной строки (QueryRow).
type Row interface {
	Scan(dest ...any) error
}

// CommandTag — результат DML-запроса (INSERT, UPDATE, DELETE).
type CommandTag interface {
	RowsAffected() int64
}
