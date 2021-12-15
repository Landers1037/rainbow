/*
Project: rainbow styles_test.go
Created: 2021/12/16 by Landers
*/

package rainbow

import (
	"testing"
)

var styleList = Colors{
	Reset,
	Bold,
	Dim,
	Italic,
	Underline,
	Blink,
	Inverse,
	Hidden,
	Strikeout,
}

func TestStyles(t *testing.T) {
	for _, c := range styleList {
		c.Println("hello world")
	}
}