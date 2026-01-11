package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gitlab.com/shaninalex/lumna/app2/models"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
)

type BoardRepository struct {
}

func NewBoardRepository() *BoardRepository {
	return &BoardRepository{}
}

var _ Repository[models.Board] = (*BoardRepository)(nil)

// Count implements Repository.
func (p *BoardRepository) Count(ctx context.Context, opts db.Expr) (int, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT count(*) FROM boards %s`,
		where,
	)

	var count int
	row := db.FromContext(ctx).QueryRowContext(ctx, query, args...)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// Create implements Repository.
func (p *BoardRepository) Create(ctx context.Context, entry *models.Board) error {
	query := `
		INSERT INTO boards (name, project_id)
		VALUES (?, ?)
	`

	result, err := db.FromContext(ctx).ExecContext(ctx, query, entry.Name, entry.ProjectId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	entry.SetId(uint(id))
	return nil
}

// Delete implements Repository.
func (p *BoardRepository) Delete(ctx context.Context, id uint) error {
	res, err := db.FromContext(ctx).ExecContext(ctx, `DELETE FROM boards WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}

// Get implements Repository.
func (p *BoardRepository) Get(ctx context.Context, id uint) (*models.Board, error) {
	board := &models.Board{}
	row := db.FromContext(ctx).QueryRow(`
		SELECT id, name, project_id, settings FROM boards WHERE id = ?
	`, id)
	err := row.Scan(&board.Id, &board.Name, &board.ProjectId, &board.Settings)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorProjectNotFound
		}
		return nil, err
	}

	return board, nil
}

// List implements Repository.
func (p *BoardRepository) List(ctx context.Context, opts db.Expr) ([]*models.Board, error) {
	where, args := db.Where(opts)
	query := fmt.Sprintf(
		`SELECT id, name, project_id, settings FROM boards %s`,
		where,
	)

	rows, err := db.FromContext(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	boards := []*models.Board{}
	for rows.Next() {
		var board models.Board
		if err := rows.Scan(
			&board.Id,
			&board.Name,
			&board.ProjectId,
			&board.Settings,
		); err != nil {
			return nil, err
		}
		boards = append(boards, &board)
	}

	return boards, nil
}

// Update implements Repository.
func (p *BoardRepository) Update(ctx context.Context, listID uint, opts db.SetExpr) error {
	setQuery, args := opts.Build()
	query := fmt.Sprintf(
		`UPDATE boards %s WHERE id = ?`,
		setQuery,
	)

	args = append(args, listID)

	res, err := db.FromContext(ctx).ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}
