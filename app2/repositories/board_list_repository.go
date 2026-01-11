package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gitlab.com/shaninalex/lumna/app2/models"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
)

type BoardListRepository struct {
}

func NewBoardListRepository() *BoardListRepository {
	return &BoardListRepository{}
}

var _ Repository[models.BoardList] = (*BoardListRepository)(nil)

// Count implements Repository.
func (p *BoardListRepository) Count(ctx context.Context, opts db.Expr) (int, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT count(*) FROM board_lists %s`,
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
func (p *BoardListRepository) Create(ctx context.Context, entry *models.BoardList) error {
	query := `
		INSERT INTO board_lists (name, board_id, list_order)
		VALUES (?, ?, ?)
	`

	result, err := db.FromContext(ctx).ExecContext(ctx, query, entry.Name, entry.BoardId, entry.Order)
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
func (p *BoardListRepository) Delete(ctx context.Context, id uint) error {
	res, err := db.FromContext(ctx).ExecContext(ctx, `DELETE FROM board_lists WHERE id = ?`, id)
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
func (p *BoardListRepository) Get(ctx context.Context, id uint) (*models.BoardList, error) {
	list := &models.BoardList{}
	row := db.FromContext(ctx).QueryRow(`
		SELECT id, name, list_order, board_id FROM board_lists WHERE id = ?
	`, id)
	err := row.Scan(&list.Id, &list.Name, &list.Order, &list.BoardId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorProjectNotFound
		}
		return nil, err
	}

	return list, nil
}

// List implements Repository.
func (p *BoardListRepository) List(ctx context.Context, opts db.Expr) ([]*models.BoardList, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT id, name, list_order, board_id FROM board_lists %s`,
		where,
	)

	rows, err := db.FromContext(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lists := []*models.BoardList{}
	for rows.Next() {
		var list models.BoardList
		if err := rows.Scan(
			&list.Id,
			&list.Name,
			&list.Order,
			&list.BoardId,
		); err != nil {
			return nil, err
		}
		lists = append(lists, &list)
	}

	return lists, nil
}

// Update implements Repository.
func (p *BoardListRepository) Update(ctx context.Context, listID uint, opts db.SetExpr) error {
	setQuery, args := opts.Build()
	query := fmt.Sprintf(
		`UPDATE board_lists %s WHERE id = ?`,
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
