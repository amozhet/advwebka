//models/models.go

package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Movies struct {
	ID             int
	Title          string
	Original_title string
	Genre          string
	Release_year   time.Time
	Runtime        time.Duration
	Synopsis       string
	Rating         float64
	Director       string
	Cast           string
	Distributor    string
	Trailer_url    string
	Poster_url     string
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Role           string
}
