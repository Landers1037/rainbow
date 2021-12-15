/*
Project: rainbow styles.go
Created: 2021/12/15 by Landers
*/

package rainbow

// 样式计算
const (
	// Reset reset emphasis style
	Reset Color = iota
	// Bold bold emphasis style
	Bold
	// Dim dim emphasis style
	Dim
	// Italic italic emphasis style
	Italic
	// Underline underline emphasis style
	Underline
	// Blink blink emphasis style
	Blink
	_
	// Inverse inverse emphasis style
	Inverse
	// Hidden hidden emphasis style
	Hidden
	// Strikeout strikeout emphasis style
	Strikeout
)