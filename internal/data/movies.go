package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
/*
Hint: If you want to use omitempty and not change the key name then you can leave it
blank in the struct tag — like this: json:",omitempty". Notice that the leading comma
is still required.
*/
