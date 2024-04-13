package gui

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
)

type MainWindow struct {
	app      fyne.App
	win      fyne.Window
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewMainWindow(app fyne.App) *MainWindow {
	w := &MainWindow{
		app:      app,
		win:      app.NewWindow("GoldWatcher"),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	w.setupUI()

	return w
}

func (w *MainWindow) Show() {
	w.win.Resize(fyne.NewSize(300, 200))
	w.win.SetFixedSize(true)
	w.win.SetMaster()

	w.win.Show()
}

func (w *MainWindow) setupUI() {
}
