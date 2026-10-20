package personal_records

import (
	"context"
	"managment/database/simple_sql"
	model "managment/models"
	"uuid"

	"github.com/jackc/pgx/v5"
)

type Managment struct {
	conn *pgx.Conn
	ctx  context.Context
}

func NewManagment(conn *pgx.Conn, ctx context.Context) *Managment {
	return &Managment{
		conn: conn,
		ctx:  ctx,
	}
}

func (m *Managment) AddWorker(worker model.Worker) error {
	if err := simple_sql.InsertRow(m.conn, m.ctx, worker); err != nil {
		return err
	}

	return nil
}

func (m *Managment) GetAllWorkers() (error, []model.Worker) {
	err, workers := simple_sql.SelectRows(m.conn, m.ctx)

	if err != nil {
		return err, nil
	}

	return nil, workers

}

func (m *Managment) DeleteWorker(id uuid.UUID) error {
	if err := simple_sql.DeleteRow(m.conn, m.ctx, id); err != nil {
		return err
	}

	return nil

}
