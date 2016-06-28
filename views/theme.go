package views

import (
  ui "github.com/nsf/termbox-go"
)

// Default Theme
var ThemeDefault Theme = Theme{
	Name: "Default", 
	Border: Border{
		// Show border or not
		Visible: true,

		// Left Border Character
		StyleLeft: ' ',

		// Left Border Charachter
		StyleRight: ' ',

		// Top Border Character
		StyleTop: ' ',

		// Bottom Border Character
		StyleBottom: ' ',

		// Top Left Charachter
		StyleTopLeft: ' ',

		// Top Right Charachter
		StyleTopRight: ' ',

		// Bottom Left Character
		StyleBottomLeft: ' ',

		// Bottom Right Carachter
		StyleBottomRight: ' ',

		// Right Border
		ColorFgRight:  ui.ColorDefault,
		ColorBgRight: ui.ColorDefault,

		// Bottom Border
		ColorFgBottom: ui.ColorDefault,
		ColorBgBottom: ui.ColorGreen,

		// Left Border
		ColorFgLeft: ui.ColorDefault,
		ColorBgLeft: ui.ColorDefault,

		// Top Border
		ColorFgTop: ui.ColorDefault,
		ColorBgTop: ui.ColorCyan,

		// Top Left Corner
		ColorFgTopLeft: ui.ColorDefault,
		ColorBgTopLeft: ui.ColorCyan,

		// Top Right Corner
		ColorFgTopRight: ui.ColorDefault,
		ColorBgTopRight: ui.ColorCyan,

		// Bottom Left Corner
		ColorFgBottomLeft: ui.ColorDefault,
		ColorBgBottomLeft: ui.ColorYellow,
		
		// Bottom Right Corner
		ColorFgBottomRight: ui.ColorDefault,
		ColorBgBottomRight: ui.ColorYellow,
	},
}

// Simple Theme
var ThemeSimple Theme = Theme{
	Name: "Simple", 
	Border: Border{
		// Show border or not
		Visible: true,

		// Left Border Character
		StyleLeft: '│',

		// Left Border Charachter
		StyleRight: '│',

		// Top Border Character
		StyleTop: '─',

		// Bottom Border Character
		StyleBottom: '─',

		// Top Left Charachter
		StyleTopLeft: '┌',

		// Top Right Charachter
		StyleTopRight: '┐',

		// Bottom Left Character
		StyleBottomLeft: '└',

		// Bottom Right Carachter
		StyleBottomRight: '┘',

		// Right Border
		ColorFgRight:  ui.ColorDefault,
		ColorBgRight: ui.ColorDefault,

		// Bottom Border
		ColorBgBottom: ui.ColorDefault,
		ColorFgBottom: ui.ColorDefault,

		// Left Border
		ColorFgLeft: ui.ColorDefault,
		ColorBgLeft: ui.ColorDefault,

		// Top Border
		ColorFgTop: ui.ColorDefault,
		ColorBgTop: ui.ColorDefault,

		// Top Left Corner
		ColorFgTopLeft: ui.ColorDefault,
		ColorBgTopLeft: ui.ColorDefault,

		// Top Right Corner
		ColorFgTopRight: ui.ColorDefault,
		ColorBgTopRight: ui.ColorDefault,

		// Bottom Left Corner
		ColorBgBottomLeft: ui.ColorDefault,
		ColorFgBottomLeft: ui.ColorDefault,

		// Bottom Right Corner
		ColorBgBottomRight: ui.ColorDefault,
		ColorFgBottomRight: ui.ColorDefault,

	},
}

// Bobbys Theme
var ThemeBobby Theme = Theme{
	Name: "bobby", 
	Border: Border{
		// Show border or not
		Visible: true,

		// Left Border Character
		StyleLeft: '|',

		// Left Border Charachter
		StyleRight: '|',

		// Top Border Character
		StyleTop: '-',

		// Bottom Border Character
		StyleBottom: '-',

		// Top Left Charachter
		StyleTopLeft: '+',

		// Top Right Charachter
		StyleTopRight: '+',

		// Bottom Left Character
		StyleBottomLeft: '+',

		// Bottom Right Carachter
		StyleBottomRight: '+',

		// Right Border
		ColorFgRight:  ui.ColorDefault,
		ColorBgRight: ui.ColorDefault,

		// Bottom Border
		ColorBgBottom: ui.ColorDefault,
		ColorFgBottom: ui.ColorDefault,

		// Left Border
		ColorFgLeft: ui.ColorDefault,
		ColorBgLeft: ui.ColorDefault,

		// Top Border
		ColorFgTop: ui.ColorDefault,
		ColorBgTop: ui.ColorDefault,

		// Top Left Corner
		ColorFgTopLeft: ui.ColorDefault,
		ColorBgTopLeft: ui.ColorDefault,

		// Top Right Corner
		ColorFgTopRight: ui.ColorDefault,
		ColorBgTopRight: ui.ColorDefault,

		// Bottom Left Corner
		ColorBgBottomLeft: ui.ColorDefault,
		ColorFgBottomLeft: ui.ColorDefault,

		// Bottom Right Corner
		ColorBgBottomRight: ui.ColorDefault,
		ColorFgBottomRight: ui.ColorDefault,

	},
}


type Theme struct {
	Border Border
	Name string
}

