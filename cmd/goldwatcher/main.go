package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/zivlakmilos/goldwatcher/gui"
)

func main() {
	a := app.NewWithID("ra.zivlak.goldwatcher")

	w := gui.NewMainWindow(a)
	w.Show()

	a.Run()
}
