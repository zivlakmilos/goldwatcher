package main

import (
	"database/sql"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/zivlakmilos/goldwatcher/private/api"
	"github.com/zivlakmilos/goldwatcher/private/gui"
	"github.com/zivlakmilos/goldwatcher/private/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	a := app.NewWithID("ra.zivlak.goldwatcher")

	db, err := connectSQL(a)
	if err != nil {
		log.Panic(err)
	}

	repo, err := setupDB(db)
	if err != nil {
		return
	}

	api.LoadCurrency(a)

	w := gui.NewMainWindow(a, repo)
	w.Show()

	a.Run()
}

func connectSQL(a fyne.App) (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = a.Storage().RootURI().Path() + "/sql.db"
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupDB(db *sql.DB) (repository.Repository, error) {
	repo := repository.NewSQLiteRepository(db)

	err := repo.Migrate()
	if err != nil {
		return nil, err
	}

	return repo, nil
}
