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
}

type View struct {
  ViewAttribute
  Title string
  FgColor ui.Attribute
  BgColor ui.Attribute
  Border bool
  BorderFgColor ui.Attribute
  BorderBgColor ui.Attribute
}

type List struct {
  ViewAttribute
  Hosts []string
}

func (l List) Draw() {
  for i, h := range l.Hosts {
    for j, c := range h {
      ui.SetCell(j+l.PosX, i+l.PosY, c, ui.ColorCyan, ui.ColorDefault)  
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
    Border: true,
    BorderFgColor: ui.ColorDefault,
    BorderBgColor: ui.ColorDefault,
  }
  v.Width = 5
  v.Height = 5
  v.PosX = 0
  v.PosY = 0
  return v
}

func (v View) NewList() List {
  l := List{
    Hosts: []string{"localhost", "127.0.0.1", "pod.amimof.com"},
  }
  l.Width = 10
  l.Height = 5
  l.PosX = v.PosX+2
  l.PosY = v.PosY+2
  return l
}
