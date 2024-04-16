package gui

import "testing"

func TestAppGetPriceText(t *testing.T) {
	open, _, _ := testMainWindow.getPriceText()
	if open.Text != "Open: $2377.0000 USD" {
		t.Error("error: wrong price returned", open.Text)
	}
}

func TestAppSetupToolBar(t *testing.T) {
	tb := testMainWindow.setupToolBar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar")
	}
}
