package simple_sql

import (
	"context"
	model "managment/models"

	"github.com/jackc/pgx/v5"
)

func InsertRow(conn *pgx.Conn, ctx context.Context, w model.Worker) error {
	sqlQuery := `
	INSERT INTO personnel_managment (id, fullname, position)
	VALUES ($1, $2, $3);
	`

	_, err := conn.Exec(ctx, sqlQuery, w.ID, w.FullName, w.Position)

	return err
}
