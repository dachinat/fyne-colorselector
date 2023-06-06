package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dachinat/colornameconv"
	"github.com/lusingander/colorpicker"
	"image/color"
)

var (
	defaultColor = color.NRGBA{0xff, 0x00, 0x00, 0xff}
	selectedType = "Hue"
)

func main() {
	a := app.New()
	w := a.NewWindow("Color Selector")

	labelText := "No color selected"

	label := widget.NewLabel(labelText)

	content := container.New(layout.NewCenterLayout(), label)

	btn := widget.NewButton("Browse Colors", func() {
		openPicker(w, label)
	})

	choices := widget.NewSelect([]string{"Hue", "Hue Circle", "Value", "Saturation"}, func(value string) {
		selectedType = value
	})

	top := container.NewGridWithColumns(2, btn, choices)

	w.SetContent(
		container.NewBorder(top, nil, nil, nil, content),
	)

	w.Resize(fyne.NewSize(400, 340))
	w.ShowAndRun()
}

func openPicker(w fyne.Window, l *widget.Label) {
	//d := dialog.NewColorPicker("Pick a color", "Browse colors", func(c color.Color) {
	//	hex := colorToHex(c)
	//	l.Text = hex + " " + HexToName(hex)
	//
	//	l.Refresh()
	//}, w)
	//
	//d.Show()

	var currentSimple color.Color
	currentSimple = defaultColor

	var styleType colorpicker.PickerStyle
	if selectedType == "Hue" {
		styleType = colorpicker.StyleHue
	} else if selectedType == "Hue Circle" {
		styleType = colorpicker.StyleHueCircle
	} else if selectedType == "Value" {
		styleType = colorpicker.StyleValue
	} else if selectedType == "Saturation" {
		styleType = colorpicker.StyleSaturation
	}

	picker := colorpicker.New(200, styleType)
	picker.SetOnChanged(func(c color.Color) {
		hex := colorToHex(c)
		l.Text = hex + " " + HexToName(hex)
		l.Refresh()
	})
	content := fyne.NewContainer(picker)

	picker.SetColor(currentSimple)
	dialog.ShowCustom("Select color", "OK", content, w)
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02X%02X%02X", r>>8, g>>8, b>>8)
}

func HexToName(hex string) string {
	name, _ := colornameconv.New(hex)
	return name
}
