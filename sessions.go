package main

import (
  ui "github.com/nsf/termbox-go"
  "github.com/amimof/sessions-go/lib"
  "fmt"
)

var ah int
var bbw, bbh int
var backbuffer []ui.Cell
var hosts []string = []string{"gotsva1012.volvocars.biz", "gotsva1014.volvocars.biz", "gotsvl2057.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl4056.got.volvocars.biz", "gotsvl4056.got.volvocars.biz"}

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

func redraw() {
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
  hostv.Height = 10
  hostv.PosX = 10
  hostv.PosY = 10

  // Draw
  hostl := hostv.NewList()
  hostl.Hosts = hosts
  //hostl.SetActive(1)
  lib.Render(main, hostv)

  ui.Flush()
}

func resizeBackBuffer(w, h int) {
  bbw, bbh = w, h
  backbuffer = make([]ui.Cell, w*h)
}

func main() {

  // Init termbox
  err := ui.Init()
  if err != nil {
    panic(err)
  }
  defer ui.Close()

  // Init views
  w, h := ui.Size()
  resizeBackBuffer(w, h)
  redraw()

  ui.SetInputMode(ui.InputEsc)
  mainloop:
  for {
    switch ev := ui.PollEvent(); ev.Type {
    case ui.EventKey:
      if ev.Key == ui.KeyEsc {
        break mainloop
      }
      if ev.Key == ui.KeyArrowDown {
        //hostl.SetActive(hostl.GetActive()+1)
      }
      if ev.Key == ui.KeyArrowUp {
        //hostl.SetActive(hostl.GetActive()-1)
      }
    case ui.EventResize:
      resizeBackBuffer(ev.Width, ev.Height)
    }
    redraw()
  }

}
