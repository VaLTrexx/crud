package repository

import (
	"context"

	"github.com/VaLTrexx/crud/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *models.Task) error {
	query := `
		INSERT INTO tasks (title, priority, position, due_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	return r.db.QueryRow(
		ctx,
		query,
		task.Title,
		task.Priority,
		task.Position,
		task.DueDate,
	).Scan(&task.ID)
}

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	query := `
		SELECT id, title, priority, position, due_date
		FROM tasks
		ORDER BY position ASC;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Priority,
			&task.Position,
			&task.DueDate,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
