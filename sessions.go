package main

import (
  ui "github.com/nsf/termbox-go"
  "github.com/amimof/sessions-go/lib"
  "fmt"
)


var body *lib.View
var ah int
var bbw, bbh int
var backbuffer []ui.Cell
var hosts []string = []string{"gotsva1012.volvocars.biz", "gotsva1014.volvocars.biz", "gotsvl2057.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl4011.got.volvocars.biz"}

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

/*func redraw() {
  ui.Clear(ui.ColorDefault, ui.ColorDefault)
  copy(ui.CellBuffer(), backbuffer)

  // Draw the workspace
  main := lib.NewView()
  main.Title = "Sessions v1.0.0b"
  main.Width = bbw
  main.Height = bbh
  main.Border.StyleLeft, main.Border.StyleRight, main.Border.StyleTop, main.Border.StyleBottom = ' ', ' ', ' ', ' '
  main.Border.ColorBgTop = ui.ColorCyan
  main.Border.ColorFgTop = ui.ColorDefault

  // Draw the view holding the host list
  hostv := lib.NewView()
  hostv.Title = "Hosts"
  hostv.Width = 50
  hostv.Height = 13
  hostv.PosX = 10
  hostv.PosY = 10
  hostv.FgColor = ui.ColorRed

  // Create a list on the hostv view
  var hostl lib.List = hostv.NewList()
  hostl.Hosts = hosts
  hostl.SetActive(1)
  hostv.List = hostl

  // Render 
  lib.Render(main, hostv)

  ui.Flush()
}
*/

func render(view... *lib.View) {
  ui.Clear(ui.ColorDefault, ui.ColorDefault)
  copy(ui.CellBuffer(), backbuffer)

  body.Width = bbw
  body.Height = bbh

  // Render
  lib.Render(view...)
  ui.Flush()

}

func resizeBackBuffer(w, h int) {
  bbw, bbh = w, h
  backbuffer = make([]ui.Cell, w*h)
}

func main() {

  // Init termbox
  err := lib.Init()
  if err != nil {
    panic(err)
  }
  defer lib.Close()

  // Draw the workspace
  body = lib.NewView()
  body.Title = "Sessions v1.0.0b"
  body.Width = bbw
  body.Height = bbh
  body.Border.StyleLeft, body.Border.StyleRight, body.Border.StyleTop, body.Border.StyleBottom = ' ', ' ', ' ', ' '
  body.Border.ColorBgTop = ui.ColorCyan
  body.Border.ColorFgTop = ui.ColorDefault

  // Draw the view holding the host list
  hostv := lib.NewView()
  hostv.Title = "Hosts"
  hostv.Width = 50
  hostv.Height = 13
  hostv.PosX = 10
  hostv.PosY = 10
  hostv.FgColor = ui.ColorRed

  // Create a list on the hostv view
  hostl := hostv.NewList()
  hostl.Hosts = hosts
  hostv.List = hostl

  // Init views
  w, h := ui.Size()
  resizeBackBuffer(w, h)

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
      }
      if ev.Key == ui.KeyArrowUp {
        hostl.Prev()
      }
    case ui.EventResize:
      resizeBackBuffer(ev.Width, ev.Height)
    }
    render(body, hostv)
  }

}
