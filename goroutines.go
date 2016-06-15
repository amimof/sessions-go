package main

import (
  ui "github.com/nsf/termbox-go"
  "fmt"
  "sync"
)

func f(s string) {
  for i := 0; i < 3; i++ {
    fmt.Println(s, ":", i)
  }
}

func Init() error {
  if err := ui.Init(); err != nil {
    return err
  }

  //go HookTermboxEvents()

  go Render("asdasd")

  return nil
}

func Close() {
  ui.Close()
}

var renderLock sync.Mutex

func Render(s string) {

  for i := 0; i < 10; i++ {
      ui.SetCell(i, i, 'x', ui.ColorDefault, ui.ColorDefault)
  }

  renderLock.Lock()
  ui.Flush()
  renderLock.Unlock()
}

func HookTermboxEvents() {
  for {
    switch ev := ui.PollEvent(); ev.Type {
    case ui.EventKey:
      if ev.Key == ui.KeyEsc {
      }
      if ev.Key == ui.KeyArrowDown {
      }
      if ev.Key == ui.KeyArrowUp {

      }
    case ui.EventResize:

    }

  }
}

func main() {

  err := Init()
  if err != nil {
    panic(err)
  }
  defer Close()

  HookTermboxEvents()

  // Synchronously
  //f("direct")

  // Async
  //go f("goroutine")

  //Render("tesdt")


}
