//models/mysql/news.go

package mysql

import (
	"AituNews/pkg/models"
	"database/sql"
	"errors"
	"time"
)

type MoviesModel struct {
	DB *sql.DB
}

func (m *MoviesModel) Get(id int) (*models.Movies, error) {

	stmt := `SELECT id, title, original_title, genre, release_year, runtime, synopsis, rating, director, cast, distributor, trailer_url, poster_url
	FROM movies
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Movies{}

	err := row.Scan(&s.ID, &s.Title, &s.Original_title, &s.Genre, &s.Release_year, &s.Runtime, &s.Synopsis, &s.Rating, &s.Director, &s.Cast, &s.Distributor, &s.Trailer_url, &s.Poster_url)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil

}

func (m *MoviesModel) Latest() ([]*models.Movies, error) {
	stmt := `SELECT id, title, original_title, genre, release_year, runtime, synopsis, rating, director, cast, distributor, trailer_url, poster_url
	FROM movies ORDER BY release_year DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	news := []*models.Movies{}
	for rows.Next() {
		s := &models.Movies{}

		err := rows.Scan(&s.ID, &s.Title, &s.Original_title, &s.Genre, &s.Release_year, &s.Runtime, &s.Synopsis, &s.Rating, &s.Director, &s.Cast, &s.Distributor, &s.Trailer_url, &s.Poster_url)
		if err != nil {
			return nil, err
		}
		news = append(news, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}

func (m *MoviesModel) Insert(title, original_title, genre string, release_year time.Time, runtime time.Duration, synopsis string, rating float64, director, cast, distributor, trailer_url, poster_url string) (int, error) {
	stmt := `INSERT INTO news ( title, original_title, genre, release_year, runtime, synopsis, rating, director, cast, distributor, trailer_url, poster_url)
    VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`

	result, err := m.DB.Exec(stmt, title, original_title, genre, release_year, runtime, synopsis, rating, director, cast, distributor, trailer_url, poster_url)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m *MoviesModel) GetMovieByGenre(genre string) ([]*models.Movies, error) {
	stmt := `
        SELECT id, title, original_title, genre, release_year, runtime, synopsis, rating, director, cast, distributor, trailer_url, poster_url
        FROM movies
        WHERE genre = ?
        ORDER BY release_year DESC
    `

	rows, err := m.DB.Query(stmt, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movies

	for rows.Next() {
		s := &models.Movies{}
		err := rows.Scan(&s.ID, &s.Title, &s.Original_title, &s.Genre, &s.Release_year, &s.Runtime, &s.Synopsis, &s.Rating, &s.Director, &s.Cast, &s.Distributor, &s.Trailer_url, &s.Poster_url)
		if err != nil {
			return nil, err
		}
		movies = append(movies, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
