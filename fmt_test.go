/*
Project: rainbow fmt_test.go
Created: 2021/12/15 by Landers
*/

package rainbow

import (
	"testing"
)

func TestPrintln(t *testing.T) {
	Println(Colors{Red, BGBlue}, "hello", "world")
}

func TestPrintf(t *testing.T) {
	Printf(Colors{Red, BGBlue}, "%s, %s", "hello", "world")
}

func TestPrint(t *testing.T) {
	Print(Colors{Red, BGBlue}, "hello", "world")
}
