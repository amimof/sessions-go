package lib

import (
  ui "github.com/nsf/termbox-go"
)

type Block interface {
  Draw()
}

type ViewAttribute struct {
  Width int
  Height int
  PosX int
  PosY int
  Border bool
}

type View struct {
  ViewAttribute
  Title string
  FgColor ui.Attribute
  BgColor ui.Attribute
  BorderFgColor ui.Attribute
  BorderBgColor ui.Attribute
}

type List struct {
  View *View
  ViewAttribute
  Hosts []string
  Active int
}

func (l List) Length() int {
  return len(l.Hosts)
}

func (l List) Next() {
  l.SetActive(l.GetActive()+1)
}

func (l List) Prev() {
  l.SetActive(l.GetActive()-1)
}


func (l List) SetActive(n int) {
  l.ClearActive()
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

  for j, c := range l.Hosts[a] {
    ui.SetCell(j+l.PosX, a+l.PosY, c, ui.ColorWhite, ui.ColorGreen)
    l.Active = a
  }
}

func (l List) GetActive() int {
  return l.Active
}

func (l List) ClearActive() {
  l.Draw()
}

func (l List) Draw() {
  for i, h := range l.Hosts {
    for j, c := range h {
      ui.SetCell(j+l.PosX, i+l.PosY, c, ui.ColorCyan, ui.ColorDefault)
    }
  }
}

func (l List) Clear() {
  for i, h := range l.Hosts {
    for j := range h {
      ui.SetCell(j+l.PosX, i+l.PosY, ' ', ui.ColorDefault, ui.ColorDefault)
    }
  }
}

func (v View) Draw() {

  // Background X
  for x := 0; x < v.Width; x++ {
    ui.SetCell(x+v.PosX, 0+v.PosX, ' ', v.BgColor, v.BgColor)
    for y := 0; y < v.Height; y++ {
      ui.SetCell(x+v.PosX, y+v.PosX, ' ', v.BgColor, v.BgColor)
    }
  }

  if v.Border {
    for i := 1; i < v.Height; i++ {
      // Left border
      ui.SetCell(0+v.PosX, i+v.PosY, '|', v.BorderFgColor, v.BorderBgColor)
      // Right border
      ui.SetCell(v.Width-1+v.PosX, i+v.PosY, '|', v.BorderFgColor, v.BorderBgColor)
      // Top left corner
      ui.SetCell(0+v.PosX, 0+v.PosY, '+', v.BorderFgColor, v.BorderBgColor)
      // Bottom left corner
      ui.SetCell(0+v.PosX, v.Height-1+v.PosY, '+', v.BorderFgColor, v.BorderBgColor)
    }
    // Width
    for i := 1; i < v.Width; i++ {
      // Top Border
      ui.SetCell(i+v.PosX, 0+v.PosY, '-', v.BorderFgColor, v.BorderBgColor)
      // Bottom Border
      ui.SetCell(i+v.PosX, v.Height-1+v.PosY, '-', v.BorderFgColor, v.BorderBgColor)
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
		ui.SetCell(i+x+v.PosX, 0+v.PosY, c, v.BorderFgColor, v.BorderBgColor)
		i++
	}

}

var v View

func NewView() View {
  v = View{
    // Width: 5,
    // Height: 5,
    // PosX: 0,
    // PosY: 0,
    FgColor: ui.ColorDefault,
    BgColor: ui.ColorDefault,
    BorderFgColor: ui.ColorDefault,
    BorderBgColor: ui.ColorDefault,
  }
  v.Border = true
  v.Width = 5
  v.Height = 5
  v.PosX = 0
  v.PosY = 0
  return v
}

func (v View) NewList() List {
  l := List{
    View: &v,
    Hosts: []string{"localhost", "127.0.0.1", "pod.amimof.com"},
  }
  l.Width = 10
  l.Height = 5
  l.PosX = v.PosX+2
  l.PosY = v.PosY+2
  return l
}
