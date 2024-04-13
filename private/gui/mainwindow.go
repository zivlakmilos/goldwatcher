package gui

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/zivlakmilos/goldwatcher/private/api"
)

type MainWindow struct {
	app            fyne.App
	win            fyne.Window
	infoLog        *log.Logger
	errorLog       *log.Logger
	priceContainer *fyne.Container
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
	openPrice, currentPrice, priceChange := w.getPriceText()

	priceContainer := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)
	w.priceContainer = priceContainer

	finalContent := container.NewVBox(priceContainer)
	w.win.SetContent(finalContent)
}

func (w *MainWindow) getPriceText() (*canvas.Text, *canvas.Text, *canvas.Text) {
	var g api.Gold
	var open, current, change *canvas.Text

	gold, err := g.GetPrices()
	if err != nil {
		gray := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
		open = canvas.NewText("Open: Unreachable", gray)
		current = canvas.NewText("Current: Unreachable", gray)
		change = canvas.NewText("Change: Unreachable", gray)
	} else {
		displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

		if gold.Price < gold.PreviousClose {
			displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255}
		}

		open = canvas.NewText(fmt.Sprintf("Open: $%.4f", gold.PreviousClose), nil)
		current = canvas.NewText(fmt.Sprintf("Current: $%.4f", gold.PreviousClose), displayColor)
		change = canvas.NewText(fmt.Sprintf("Change: $%.4f", gold.Change), displayColor)
	}

	open.Alignment = fyne.TextAlignLeading
	current.Alignment = fyne.TextAlignCenter
	change.Alignment = fyne.TextAlignTrailing

	return open, current, change
}
