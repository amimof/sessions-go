package lib

import (
  ui "github.com/nsf/termbox-go"
  "fmt"
)

type Block interface {
  Draw()
}

type Border struct {
  Visible bool
  StyleLeft rune
  StyleRight rune
  StyleTop rune
  StyleBottom rune
  ColorFgLeft ui.Attribute
  ColorFgRight ui.Attribute
  ColorFgTop ui.Attribute
  ColorFgBottom ui.Attribute
  ColorBgLeft ui.Attribute
  ColorBgRight ui.Attribute
  ColorBgTop ui.Attribute
  ColorBgBottom ui.Attribute
}

type ViewAttribute struct {
  Width int
  Height int
  PosX int
  PosY int
}

type View struct {
  Border Border
  ViewAttribute
  Title string
  FgColor ui.Attribute
  BgColor ui.Attribute
  BorderFgColor ui.Attribute
  BorderBgColor ui.Attribute
  List List
}

type List struct {
  View *View
  ViewAttribute
  Hosts []string
  Active int
}

func (l *List) Length() int {
  return len(l.Hosts)
}

func (l *List) Next() {
  l.SetActive(l.GetActive()+1)
}

func (l *List) Prev() {
  l.SetActive(l.GetActive()-1)
}


func (l *List) SetActive(n int) {
  //l.ClearActive()
  a := n
  // Select last index if out of range
  if n > l.Length()-1 {
     a = l.Length()-1
  } else {
    a = n
  }
  // Select first index if out of range
  if n < 0 {
    a = 0
  } else {
    a = n
  }

  fmt.Println(a)
  fmt.Println(l.Hosts)

  // for j, c := range l.Hosts[0] {
  //   ui.SetCell(j+l.PosX, a+l.PosY, c, ui.ColorWhite, ui.ColorGreen)
  //   l.Active = a
  // }
}

func (l *List) GetActive() int {
  return l.Active
}

func (l *List) ClearActive() {
  l.Draw()
}

func (l *List) Clear() {
  for i, h := range l.Hosts {
    for j := range h {
      ui.SetCell(j+l.PosX, i+l.PosY, ' ', ui.ColorDefault, ui.ColorDefault)
    }
  }
}

func (l *List) Draw() {
  for i, h := range l.Hosts {
    for j, c := range h {
      ui.SetCell(j+l.PosX, i+l.PosY, c, ui.ColorCyan, ui.ColorDefault)
    }
  }
  //l.SetActive(1)
}

func (v *View) Draw() {

  // Background X
  for x := 0; x < v.Width; x++ {
    ui.SetCell(x+v.PosX, 0+v.PosX, ' ', v.BgColor, v.BgColor)
    for y := 0; y < v.Height; y++ {
      ui.SetCell(x+v.PosX, y+v.PosX, ' ', v.BgColor, v.BgColor)
    }
  }

  // Draw list
  v.List.Draw()

  if v.Border.Visible {
    for i := 1; i < v.Height; i++ {
      // Left border
      ui.SetCell(0+v.PosX, i+v.PosY, v.Border.StyleLeft, v.Border.ColorFgLeft, v.Border.ColorBgLeft)
      // Right border
      ui.SetCell(v.Width-1+v.PosX, i+v.PosY, v.Border.StyleRight, v.Border.ColorFgRight, v.Border.ColorBgRight)
      // Top left corner
      ui.SetCell(0+v.PosX, 0+v.PosY, '+', v.BorderFgColor, v.BorderBgColor)
      // Bottom left corner
      ui.SetCell(0+v.PosX, v.Height-1+v.PosY, '+', v.BorderFgColor, v.BorderBgColor)
    }
    // Width
    for i := 1; i < v.Width; i++ {
      // Top Border
      ui.SetCell(i+v.PosX, 0+v.PosY, v.Border.StyleTop, v.Border.ColorFgTop, v.Border.ColorBgTop)
      // Bottom Border
      ui.SetCell(i+v.PosX, v.Height-1+v.PosY, v.Border.StyleBottom, v.Border.ColorFgBottom, v.Border.ColorBgBottom)
      // Bottom right corner
      ui.SetCell(v.Width-1+v.PosX, v.Height-1+v.PosY, '+', v.BorderFgColor, v.BorderBgColor)
      // Top right corner
      ui.SetCell(v.Width-1+v.PosX, 0+v.PosY, '+', v.BorderFgColor, v.BorderBgColor)
    }
  }

  // Title
  tlen := len(v.Title)
  x := (v.Width/2)-(tlen/2)
  for i, c := range v.Title {
		ui.SetCell(i+x+v.PosX, 0+v.PosY, c, v.Border.ColorFgTop, v.Border.ColorBgTop)
		i++
	}

}

// Call Draw method on all views
func Render(v... *View) {
  for _, view := range v {
    view.Draw()
  }
}

func NewView() *View {
  v := &View{
    FgColor: ui.ColorDefault,
    BgColor: ui.ColorDefault,
    BorderFgColor: ui.ColorDefault,
    BorderBgColor: ui.ColorDefault,
  }
  v.Border.Visible = true

  // Border style left/right
  v.Border.StyleLeft, v.Border.StyleRight = '|', '|'
  v.Border.StyleTop, v.Border.StyleBottom = '-', '-'

  // Border colors
  v.Border.ColorFgLeft = ui.ColorDefault
  v.Border.ColorFgRight = ui.ColorDefault
  v.Border.ColorFgTop = ui.ColorDefault
  v.Border.ColorFgBottom = ui.ColorDefault
  v.Border.ColorBgLeft = ui.ColorDefault
  v.Border.ColorBgRight = ui.ColorDefault
  v.Border.ColorBgTop = ui.ColorDefault
  v.Border.ColorBgBottom = ui.ColorDefault

  // Size
  v.Width = 5
  v.Height = 5

  // Position
  v.PosX = 0
  v.PosY = 0
  return v
}

func (v *View) NewList() List {
  l := List{
    Hosts: []string{},
  }
  l.Width = 10
  l.Height = 5
  l.PosX = v.PosX+2
  l.PosY = v.PosY+2
  l.Active = 1
  v.List = l
  return v.List
}
