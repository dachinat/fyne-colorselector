package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"github.com/dachinat/colornameconv"
	"github.com/lusingander/colorpicker"
	"image/color"
)

var (
	defaultColor  = color.NRGBA{0xff, 0x00, 0x00, 0xff}
	selectedType  = "Hue"
	selectedHex   = ""
	currentSimple color.Color
)

func main() {
	a := app.New()
	w := a.NewWindow("Color Selector")

	labelText := "No color selected"

	label := widget.NewLabel(labelText)

	btn2 := widget.NewButtonWithIcon("Copy Hex", theme.ContentCopyIcon(), func() {
		clipboard.WriteAll(selectedHex)
	})

	contentLabel := container.New(layout.NewCenterLayout(), label)
	contentBtn := container.New(layout.NewCenterLayout(), btn2)
	contentRect := canvas.NewRectangle(defaultColor)
	contentRect.SetMinSize(fyne.NewSize(75, 75))
	content := container.NewBorder(contentLabel, contentBtn, nil, nil, container.NewCenter(contentRect))

	btn := widget.NewButton("Browse Colors", func() {
		openPicker(w, label, contentRect)
	})

	choices := widget.NewSelect([]string{"Hue", "Hue Circle", "Value", "Saturation"}, func(value string) {
		selectedType = value
	})

	choices.SetSelected("Hue")

	top := container.NewGridWithColumns(2, btn, choices)

	w.SetContent(
		container.NewBorder(top, nil, nil, nil, content),
	)

	w.Resize(fyne.NewSize(400, 340))
	w.ShowAndRun()
}

func openPicker(w fyne.Window, l *widget.Label, r *canvas.Rectangle) {
	//d := dialog.NewColorPicker("Pick a color", "Browse colors", func(c color.Color) {
	//	hex := colorToHex(c)
	//	l.Text = hex + " " + HexToName(hex)
	//
	//	l.Refresh()
	//}, w)
	//
	//d.Show()

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

	if currentSimple != nil {
		picker.SetColor(currentSimple)
	} else {
		picker.SetColor(defaultColor)
	}

	picker.SetOnChanged(func(c color.Color) {
		hex := colorToHex(c)
		l.Text = hex + " (" + HexToName(hex) + ")"
		l.Refresh()

		r.FillColor = c
		r.Refresh()

		selectedHex = hex

		currentSimple = c
	})
	content := fyne.NewContainer(picker)

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
