-- name: createFavMovies :one
INSERT INTO fav_movies (
  id ,
  user_id ,
  movies
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetFavMoviesByUserId :one
SELECT * FROM fav_movies
WHERE user_id = $1 LIMIT 1;

-- name: UpdateFavMoviesByUserId :one
UPDATE fav_movies 
SET movies = $2
WHERE user_id = $1
RETURNING *;
