package views

import (
  ui "github.com/nsf/termbox-go"
  //"fmt"
)

type Block interface {
  Draw()
}

type ViewAttribute struct {
  
  // Size
  Width int
  Height int
  
  // Position
  PosX int
  PosY int

}

type View struct {
  Border Border
  Theme Theme
  ViewAttribute
  Title string
  FgColor ui.Attribute
  BgColor ui.Attribute
  List *List
  TextBox *TextBox
}

type List struct {
  ViewAttribute
  Hosts []string
  Active int
  View *View
}

type TextBox struct {
  Text string
  ViewAttribute
  AutoWidth bool
  View *View
}

/**
 *
 * PRIVATE METHODS
 * 
 */


// DEPRECATED
func cut(s string, n int) string {
  return s
}

/**
 *
 * PUBLIC METHODS
 * 
 */

func (l *List) Length() int {
  return len(l.Hosts)
}

func (l *List) Next() {
  l.SetActive(l.Active+1)
}

func (l *List) Prev() {
  l.SetActive(l.Active-1)
}


func (l *List) SetActive(n int) {

  //l.ClearActive()
  a := n
  // Select last index if out of range
  if n >= len(l.Hosts)-1 {
     a = len(l.Hosts)-1
  } else {
    a = n
  }
  // Select first index if out of range
  if n < 0 {
    a = 0
  } else {
    a = n
  }


  host := l.Hosts[a]
  if (len(host)-2) >= l.View.Width {
    host = host[0:(l.View.Width)-7]+"..."  
  }
  for j, c := range host {
    ui.SetCell(j+l.PosX, a+l.PosY, c, ui.ColorWhite, ui.ColorGreen)
  }
  l.Active = a

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
    
    // Cut end of string if host name exeeds view width. Add "..." to end of string
    if (len(h)+2) >= l.View.Width {
      h = h[0:(l.View.Width)-7]+"..."
    }
    
    for j, c := range h {
      ui.SetCell(j+l.PosX, i+l.PosY, c, ui.ColorCyan, ui.ColorDefault)
    }

  }
  l.SetActive(l.Active)
}

func (t *TextBox) Draw() {

  text := t.Text

  // Compensate for string that are wider than the actual view. Increase the view width if so. 
  // But only do so if AutoWidth if true
  if (len(t.Text)+2) >= t.View.Width {
    if t.AutoWidth {
      t.View.Width = len(t.Text)+4
    } else {
      // If AutoWidth is false, cut the text that doesn't fit into the view
      text = t.Text[0:(t.View.Width-7)]+"..."
    }
  } 

  for i, c := range text {
    ui.SetCell(i+t.PosX, t.PosY, c, ui.ColorDefault, ui.ColorDefault)
  }
}

func (t *TextBox) SetText(str string) {
  t.Text = str
  //t.Draw()
}

func (v *View) Draw() {

  // Background X
  for x := 0; x < v.Width; x++ {
    ui.SetCell(x+v.PosX, 0+v.PosX, ' ', v.BgColor, v.BgColor)
    for y := 0; y < v.Height; y++ {
      ui.SetCell(x+v.PosX, y+v.PosX, ' ', v.BgColor, v.BgColor)
    }
  }

  // Draw list if any on the view
  if v.List != nil {
    // Increase height of view to number of hosts
    v.Height = len(v.List.Hosts)+2
    v.List.Draw()
  }

  // Draw text box if any on the view
  if v.TextBox != nil {
    v.TextBox.Draw()
  }
  

  if v.Theme.Border.Visible {
    for i := 1; i < v.Height; i++ {
      // Left border
      ui.SetCell(0+v.PosX, i+v.PosY, v.Theme.Border.StyleLeft, v.Theme.Border.ColorFgLeft, v.Theme.Border.ColorBgLeft)
      // Right border
      ui.SetCell(v.Width-1+v.PosX, i+v.PosY, v.Theme.Border.StyleRight, v.Theme.Border.ColorFgRight, v.Theme.Border.ColorBgRight)
      // Top left corner
      ui.SetCell(0+v.PosX, 0+v.PosY, v.Theme.Border.StyleTopLeft, v.Theme.Border.ColorFgTopLeft, v.Theme.Border.ColorBgTopLeft)
      // Bottom left corner
      ui.SetCell(0+v.PosX, v.Height-1+v.PosY, v.Theme.Border.StyleBottomLeft, v.Theme.Border.ColorFgBottomLeft, v.Theme.Border.ColorBgBottomLeft)
    }
    // Width
    for i := 1; i < v.Width; i++ {
      // Top Border
      ui.SetCell(i+v.PosX, 0+v.PosY, v.Theme.Border.StyleTop, v.Theme.Border.ColorFgTop, v.Theme.Border.ColorBgTop)
      // Bottom Border
      ui.SetCell(i+v.PosX, v.Height-1+v.PosY, v.Theme.Border.StyleBottom, v.Theme.Border.ColorFgBottom, v.Theme.Border.ColorBgBottom)
      // Bottom right corner
      ui.SetCell(v.Width-1+v.PosX, v.Height-1+v.PosY, v.Theme.Border.StyleBottomRight, v.Theme.Border.ColorFgBottomRight, v.Theme.Border.ColorBgBottomRight)
      // Top right corner
      ui.SetCell(v.Width-1+v.PosX, 0+v.PosY, v.Theme.Border.StyleTopRight, v.Theme.Border.ColorFgTopRight, v.Theme.Border.ColorBgTopRight)
    }
  }

  // Title
  tlen := len(v.Title)
  x := (v.Width/2)-(tlen/2)
  for i, c := range v.Title {
		ui.SetCell(i+x+v.PosX, 0+v.PosY, c, v.Theme.Border.ColorFgTop, v.Theme.Border.ColorBgTop)
		i++
	}

}

func Init() error {
  if err := ui.Init(); err != nil {
    return err
  }
  return nil
}

func Close() {
  ui.Close()
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
    Theme: ThemeDefault,
  }

  // Size
  v.Width = 5
  v.Height = 5

  // Position
  v.PosX = 0
  v.PosY = 0
  return v
}

func (v *View) NewList() *List {
  l := &List{
    Hosts: []string{},
  }
  l.Width = 10
  l.Height = 5
  l.PosX = v.PosX+2
  l.PosY = v.PosY+1
  l.Active = 0
  l.View = v
  v.List = l
  return l
}

func (v *View) NewTextBox() *TextBox {
  t := &TextBox{
    Text: "",
    AutoWidth: true,
  }
  t.PosX = v.PosX+2
  t.PosY = v.PosY+1
  t.View = v
  v.TextBox = t
  return t
}
