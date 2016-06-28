package main

import (
  ui "github.com/nsf/termbox-go"
  "github.com/amimof/sessions-go/views"
  "fmt"
)


var body *views.View
var ah int
var bbw, bbh int
var backbuffer []ui.Cell
var hosts []string = []string{"gotsva1012.volvocars.biz", "gotsva1014.volvocars.biz", "gotsvl2057.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl2121.got.volvocars.biz", "gotsva1012.volvocars.biz", "hybtestl001.got.volvocars.biz", "server01-gbg-swe.got.awesomedomain.com"}

func print_tb(x, y int, fg, bg ui.Attribute, msg string) {
	for _, c := range msg {
		ui.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printf_tb(x, y int, fg, bg ui.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}

func render(view... *views.View) {
  ui.Clear(ui.ColorDefault, ui.ColorDefault)
  copy(ui.CellBuffer(), backbuffer)

  body.Width = bbw
  body.Height = bbh

  // Render
  views.Render(view...)
  ui.Flush()

}

func resizeBackBuffer(w, h int) {
  bbw, bbh = w, h
  backbuffer = make([]ui.Cell, w*h)
}

func main() {

  // Init termbox
  err := views.Init()
  if err != nil {
    panic(err)
  }
  defer views.Close()

  // Draw the workspace
  body = views.NewView()
  body.Title = "Sessions v1.0.0b"
  body.Width = bbw
  body.Height = bbh
  body.Border.StyleLeft, body.Border.StyleRight, body.Border.StyleTop, body.Border.StyleBottom = ' ', ' ', ' ', ' '
  body.Border.ColorBgTop = ui.ColorCyan
  body.Border.ColorFgTop = ui.ColorDefault

  // Draw the view holding the host list
  hostv := views.NewView()
  hostv.Title = "Hosts"
  hostv.Width = 35
  hostv.PosX = 1
  hostv.PosY = 1
  hostv.Theme = views.ThemeSimple
  hostv.FgColor = ui.ColorRed

  // Create a list on the hostv view
  hostl := hostv.NewList()
  hostl.Hosts = hosts
  hostv.List = hostl

  // Draw the textbox view
  textv := views.NewView()
  textv.Title = "Current Host"
  textv.Width = 25
  textv.Height = 3
  textv.PosX = hostv.Width + 2
  textv.PosY = hostv.PosY
  textv.Theme = views.ThemeSimple
  
  // Draw the box 
  textb := textv.NewTextBox()
  textb.Text = "Hello World!"
  

  // Init views
  w, h := ui.Size()
  resizeBackBuffer(w, h)
  render(body, hostv, textv)

  ui.SetInputMode(ui.InputEsc)
  mainloop:
  for {
    switch ev := ui.PollEvent(); ev.Type {
    case ui.EventKey:
      if ev.Key == ui.KeyEsc {
        break mainloop
      }
      if ev.Key == ui.KeyArrowDown {
        hostl.Next()
        textb.SetText(hostl.Hosts[hostl.Active])
      }
      if ev.Key == ui.KeyArrowUp {
        hostl.Prev()
        textb.SetText(hostl.Hosts[hostl.Active])
      }
    case ui.EventResize:
      resizeBackBuffer(ev.Width, ev.Height)
    }
    render(body, hostv, textv)
  }

}
