package main

import "github.com/jackc/pgx/v5/pgxpool"

var db *pgxpool.Pool

type blogItem struct {
	ID       int    `bson:"id,omitempty" json:"id,omitempty"`
	AuthorID int    `bson:"author_id" json:"author_id"`
	Content  string `bson:"content" json:"content"`
	Title    string `bson:"title" json:"title"`
}
