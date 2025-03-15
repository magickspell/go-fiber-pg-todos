package migrations

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const createTasksTable = `
CREATE TYPE task_status AS ENUM ('new', 'in_progress', 'done');

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status task_status DEFAULT 'new',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
`

func Migrate(conn *sql.DB) error {
	_, err := conn.Exec(createTasksTable)
	return err
}
