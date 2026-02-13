package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/Endropr/ai-programming-mentor/internal/domain"
)

// PostgresRepo будет отвечать за связь Go с твоей базой ai_mentor_db
type PostgresRepo struct {
	conn *pgx.Conn
}

func NewPostgresRepo(conn *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{conn: conn}
}

// SaveMessage автоматически выполнит INSERT, который ты до этого делал вручную
func (r *PostgresRepo) SaveMessage(ctx context.Context, msg domain.Message) error {
	query := `INSERT INTO messages (user_id, role, content) VALUES ($1, $2, $3)`
	_, err := r.conn.Exec(ctx, query, msg.UserID, msg.Role, msg.Content)
	return err
}