package simple_sql

import (
	"context"
	"uuid"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(conn *pgx.Conn, ctx context.Context, id uuid.UUID) error {
	sqlQuery := `
	DELETE FROM personnel_managment
	WHERE id = $1;
	`

	_, err := conn.Exec(ctx, sqlQuery, id)

	return err
}
