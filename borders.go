package main

type borderIndex int

const (
	borderTopLeft borderIndex = iota
	borderTop
	borderTopRight

	borderLeft
	borderMiddle
	borderRight

	borderBottomLeft
	borderBottom
	borderBottomRight

	borderLine
)

type borderStyle [10]string

var borderStyles = map[string]borderStyle{
	"say": {
		" ", "_", " ",
		"|", " ", "|",
		" ", "─", " ",
		"\\",
	},
	"classicish": {
		" ", "_", " ",
		"<", " ", ">",
		" ", "-", " ",
		"\\",
	},
	"think": {
		" ", "_", " ",
		"(", " ", ")",
		" ", "─", " ",
		"o",
	},
	"unicode": {
		"┌", "─", "┐",
		"│", " ", "│",
		"└", "─", "┘",
		"╲",
	},
	"thick": {
		"┏", "━", "┓",
		"┃", " ", "┃",
		"┗", "━", "┛",
		"╲",
	},
	"rounded": {
		"╭", "─", "╮",
		"│", " ", "│",
		"╰", "─", "╯",
		"╲",
	},
}
