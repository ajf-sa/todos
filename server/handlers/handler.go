package handlers

import "github.com/alfuhigi/todos/server/db"

type Handlers struct {
	Repo *db.Repo
}

func NewHandlers(repo *db.Repo) *Handlers {
	return &Handlers{Repo: repo}
}
