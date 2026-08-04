package simple_sql

//todo:создать новую запись в таблице
//добавить ограничение уникальности заголовка(создается при создании таблицы)
import (
	"context"

	"github.com/jackc/pgx/v5"
)

//схема запроса пользователя: http->body->struct->args

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
) error {
	sqlQuery := `INSERT INTO tasks (title, description, completed, created_at) 
	VALUES ($1, $2, $3, $4)`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt)
	return err
}
