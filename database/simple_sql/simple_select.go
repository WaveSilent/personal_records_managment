package simple_sql

import (
	"context"
	model "managment/models"

	"github.com/jackc/pgx/v5"
)

func SelectRows(conn *pgx.Conn, ctx context.Context) (error, []model.Worker) {
	sqlQuery := `
		SELECT * FROM personnel_managment
	`

	rows, err := conn.Query(ctx, sqlQuery)

	if err != nil {
		return err, nil
	}

	defer rows.Close()

	workers := make([]model.Worker, 0)

	for rows.Next() {
		var worker model.Worker

		err := rows.Scan(
			&worker.ID,
			&worker.FullName,
			&worker.Position,
		)

		if err != nil {
			return err, nil
		}

		workers = append(workers, worker)
	}

	return nil, workers
}
