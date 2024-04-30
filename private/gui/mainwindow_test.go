package gui

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

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

func TestAppGetHoldingSlice(t *testing.T) {
	slice := testMainWindow.getHoldingSlice()

	if len(slice) != 3 {
		t.Error("wrong number of rows, expected 3 but got", len(slice))
	}
}

func TestAddHoldingsDialog(t *testing.T) {
	testMainWindow.addHoldingsDialog()

	test.Type(testMainWindow.addHoldingsPurchaseAmountEntry, "1")
	test.Type(testMainWindow.addHoldingsPurchasePriceEntry, "1000")
	test.Type(testMainWindow.addHoldingsPurchaseDateEntry, "2024-04-30")

	if testMainWindow.addHoldingsPurchaseDateEntry.Text != "2024-04-30" {
		t.Error("date not correct")
	}

	if testMainWindow.addHoldingsPurchaseAmountEntry.Text != "1" {
		t.Error("amount not correct")
	}

	if testMainWindow.addHoldingsPurchasePriceEntry.Text != "1000" {
		t.Error("price not correct")
	}
}
