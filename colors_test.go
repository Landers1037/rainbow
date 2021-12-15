/*
Project: rainbow colors_test.go
Created: 2021/12/15 by Landers
*/

package rainbow

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	t.Logf("\x1b[%dm%s\x1b[0m\n", Red, "hello")
}

func TestColor_Raw(t *testing.T) {
	t.Log(Black.Raw())
}

func TestColor_Fmt(t *testing.T) {
	t.Logf(Red.Fmt(), "hello")
	t.Logf(BGRed.Fmt(), "hello")
}

func TestColor_FmtStyle(t *testing.T) {
	t.Logf(Red.FmtStyle(Italic), "hello")
}

func TestColor_Println(t *testing.T) {
	Red.Println("hello world")
}

func TestColor_Printf(t *testing.T) {
	Red.Printf("name: %s", "job")
}

func TestColor_Print(t *testing.T) {
	Red.Print("hello", "world")
}

func TestResetColor(t *testing.T) {
	SetColor(Blue)
	fmt.Println("hello")
	fmt.Println("world")
	ResetColor()
}