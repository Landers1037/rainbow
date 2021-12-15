/*
Project: rainbow color_all_test.go
Created: 2021/12/15 by Landers
*/

package rainbow

import (
	"testing"
)

const (
	Text = "========"
	Data = "        "
)

var colorList = Colors{
	Black,
	Red,
	Green,
	Yellow,
	Blue,
	Magenta,
	Cyan,
	White,
	Gray,
}

var colorList2 = Colors{
	BGBlack,
	BGRed,
	BGGreen,
	BGYellow,
	BGBlue,
	BGMagenta,
	BGCyan,
	BGWhite,
	BGGray,
}

func TestColorAll(t *testing.T) {
	for _, c := range colorList {
		c.Println(Text)
	}

	for _, c := range colorList2 {
		c.Println(Data)
	}
}


