package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/Endropr/ai-programming-mentor/internal/domain"
)

type PostgresRepo struct {
	conn *pgx.Conn
}

func NewPostgresRepo(conn *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{conn: conn}
}

func (r *PostgresRepo) SaveMessage(ctx context.Context, msg domain.Message) error {
	query := `INSERT INTO messages (user_id, role, content, selected_language) VALUES ($1, $2, $3, $4)`
	
	_, err := r.conn.Exec(ctx, query, msg.UserID, msg.Role, msg.Content, msg.SelectedLanguage)
	return err
}
