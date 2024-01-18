// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: activity.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createActivity = `-- name: CreateActivity :one
INSERT INTO activities (
  id ,
  user_id ,
  view_page ,
  duration ,
  page_visited_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, user_id, view_page, duration, page_visited_at
`

type CreateActivityParams struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	ViewPage      string    `json:"view_page"`
	Duration      int32     `json:"duration"`
	PageVisitedAt time.Time `json:"page_visited_at"`
}

func (q *Queries) CreateActivity(ctx context.Context, arg CreateActivityParams) (Activity, error) {
	row := q.db.QueryRow(ctx, createActivity,
		arg.ID,
		arg.UserID,
		arg.ViewPage,
		arg.Duration,
		arg.PageVisitedAt,
	)
	var i Activity
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ViewPage,
		&i.Duration,
		&i.PageVisitedAt,
	)
	return i, err
}

const getActivitiesByUserId = `-- name: GetActivitiesByUserId :many
SELECT id, user_id, view_page, duration, page_visited_at FROM activities
WHERE user_id = $1
ORDER BY page_visited_at
`

func (q *Queries) GetActivitiesByUserId(ctx context.Context, userID uuid.UUID) ([]Activity, error) {
	rows, err := q.db.Query(ctx, getActivitiesByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Activity{}
	for rows.Next() {
		var i Activity
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ViewPage,
			&i.Duration,
			&i.PageVisitedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}