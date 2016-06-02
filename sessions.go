package main

import (
  ui "github.com/nsf/termbox-go"
  "github.com/amimof/sessions-go/lib"
  "fmt"
)

var bbw, bbh int
var backbuffer []ui.Cell
var hosts []string = []string{"gotsva1012.volvocars.biz", "gotsva1014.volvocars.biz", "gotsvl2057.volvocars.biz", "gotsvl4056.got.volvocars.biz"}

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

func redrawWidgets() {

  // Draw the workspace
  view := lib.NewView()
  view.Title = "Sessions v1.0.0b"
  view.Width = bbw
  view.Height = bbh
  view.Draw()

  list := view.NewList()
  list.Draw()

  // hostv := util.NewView()
  // hostv.Title = "Hosts"
  // hostv.Width = 50
  // hostv.Height = 10
  // hostv.PosX = 10
  // hostv.PosY = 10


  // Host list
  // for index, host := range hosts {
  //   printf_tb(10, 10+index, ui.ColorDefault, ui.ColorYellow, " %v ", host)
  // }

  printf_tb(0, bbh-1, ui.ColorDefault, ui.ColorCyan, "ESC: quit | Shift+H: Help | Arrows : Select")
}

func redraw(mx, my int) {
  ui.Clear(ui.ColorDefault, ui.ColorDefault)
  copy(ui.CellBuffer(), backbuffer)
  redrawWidgets()
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

  ui.SetInputMode(ui.InputEsc)
  mainloop:
  for {
    mx, my := -1, -1
    switch ev := ui.PollEvent(); ev.Type {
    case ui.EventKey:
      if ev.Key == ui.KeyEsc {
        break mainloop
      }
    case ui.EventResize:
      resizeBackBuffer(ev.Width, ev.Height)
    }
    redraw(mx, my)
  }

}
