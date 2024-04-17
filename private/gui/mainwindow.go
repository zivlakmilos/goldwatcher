package gui

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zivlakmilos/goldwatcher/private/api"
	"github.com/zivlakmilos/goldwatcher/resources"
)

type MainWindow struct {
	app                 fyne.App
	win                 fyne.Window
	infoLog             *log.Logger
	errorLog            *log.Logger
	priceContainer      *fyne.Container
	httpClient          *http.Client
	toolBar             *widget.Toolbar
	priceChartContainer *fyne.Container
}

func NewMainWindow(app fyne.App) *MainWindow {
	w := &MainWindow{
		app:        app,
		win:        app.NewWindow("GoldWatcher"),
		infoLog:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog:   log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		httpClient: &http.Client{},
	}

	w.setupUI()

	return w
}

func (w *MainWindow) Show() {
	w.win.Resize(fyne.NewSize(770, 410))
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

	toolBar := w.setupToolBar()
	w.toolBar = toolBar

	priceTabContent := w.setupPricesTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), canvas.NewText("Holdings contet goes here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	finalContent := container.NewVBox(priceContainer, toolBar, tabs)
	w.win.SetContent(finalContent)
}

func (w *MainWindow) getPriceText() (*canvas.Text, *canvas.Text, *canvas.Text) {
	var g api.Gold
	var open, current, change *canvas.Text

	g.Client = w.httpClient

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

		open = canvas.NewText(fmt.Sprintf("Open: $%.4f %s", gold.PreviousClose, api.Currency), nil)
		current = canvas.NewText(fmt.Sprintf("Current: $%.4f %s", gold.PreviousClose, api.Currency), displayColor)
		change = canvas.NewText(fmt.Sprintf("Change: $%.4f %s", gold.Change, api.Currency), displayColor)
	}

	open.Alignment = fyne.TextAlignLeading
	current.Alignment = fyne.TextAlignCenter
	change.Alignment = fyne.TextAlignTrailing

	return open, current, change
}

func (w *MainWindow) setupToolBar() *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return toolBar
}

func (w *MainWindow) setupPricesTab() *fyne.Container {
	chart := w.getChart()
	chartContainer := container.NewVBox(chart)
	w.priceChartContainer = chartContainer
	return chartContainer
}

func (w *MainWindow) getChart() *canvas.Image {
	apiUrl := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(api.Currency))
	var img *canvas.Image

	err := w.downloadFile(apiUrl, "gold.png")
	if err != nil {
		img = canvas.NewImageFromResource(resources.UnreachablePng)
	} else {
		img = canvas.NewImageFromFile("gold.png")
	}

	img.SetMinSize(fyne.NewSize(770, 410))

	return img
}

func (w *MainWindow) downloadFile(url, fileName string) error {
	res, err := w.httpClient.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("received wrong response code when downloading image")
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
