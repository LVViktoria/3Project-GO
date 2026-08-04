package simple_connection

//todo:создать подключение к бд

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {

	return pgx.Connect(ctx, "postgres://postgres:7576@localhost:5432/postgres")

}
