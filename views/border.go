package views

import(
	ui "github.com/nsf/termbox-go"
)

type Border struct {
	Visible bool
	StyleLeft rune
	StyleRight rune
	StyleTop rune
	StyleBottom rune
	StyleTopLeft rune
	StyleTopRight rune
	StyleBottomLeft rune
	StyleBottomRight rune
	ColorFgLeft ui.Attribute
	ColorFgRight ui.Attribute
	ColorFgTop ui.Attribute
	ColorFgBottom ui.Attribute
	ColorBgLeft ui.Attribute
	ColorBgRight ui.Attribute
	ColorBgTop ui.Attribute
	ColorBgBottom ui.Attribute
	ColorFgTopLeft ui.Attribute
	ColorBgTopLeft ui.Attribute
	ColorBgTopRight ui.Attribute
	ColorFgTopRight ui.Attribute
	ColorBgBottomLeft ui.Attribute
	ColorFgBottomLeft ui.Attribute
	ColorBgBottomRight ui.Attribute
	ColorFgBottomRight ui.Attribute
}