// app.go

package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App exposes refs to the router and database for the application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize connects to the database and create wiring to appropriate routes
func (a *App) Initialize(user, password, dbname string) {}

// Run starts the App
func (a *App) Run(addr string) {}
