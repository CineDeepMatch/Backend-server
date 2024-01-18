-- name: CreateActivity :one
INSERT INTO activities (
  id ,
  user_id ,
  view_page ,
  duration ,
  page_visited_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetActivitiesByUserId :many
SELECT * FROM activities
WHERE user_id = $1
ORDER BY page_visited_at;