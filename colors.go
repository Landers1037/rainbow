/*
Project: rainbow colors.go
Created: 2021/12/15 by Landers
*/

package rainbow

import (
	"bytes"
	"fmt"
	"strconv"
)

type Color int

// Raw 返回颜色对应的数字
func (color Color) Raw() int {
	return int(color)
}

// String 返回颜色字符串
func (color Color) String() string {
	return strconv.Itoa(color.Raw())
}

// Fmt 生成单独的无样式颜色格式串
func (color Color) Fmt() string {
	b := new(bytes.Buffer)
	b.WriteString("\x1b[")
	b.WriteString(color.String())
	b.WriteString("m")
	return fmt.Sprintf("%s%%v\x1b[0m", b.String())
}

// FmtStyle 生成带样式的格式串
// 样式见calcType
func (color Color) FmtStyle(style ...Color) string {
	b := new(bytes.Buffer)
	b.WriteString("\x1b[")
	// append styles
	for _, s := range style {
		b.WriteString(s.String())
		b.WriteString(";")
	}
	b.WriteString(color.String())
	b.WriteString("m")
	return fmt.Sprintf("%s%%v\x1b[0m", b.String())
}

func (color Color) Printf(f string, args ...interface{}) {
	fmt.Printf(
		fmt.Sprintf(color.Fmt(), fmt.Sprintf(f, args...)))
}

func (color Color) Println(args ...interface{}) {
	fmt.Println(fmt.Sprintf(color.Fmt(), fmt.Sprint(args...)))
}

func (color Color) Print(args ...interface{}) {
	fmt.Print(fmt.Sprintf(color.Fmt(), fmt.Sprint(args...)))
}

// 前景色
const (
	Black Color = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Gray
)

// 背景色
const (
	BGBlack Color = 40 + iota
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
	BGGray
)

// ResetColor 重置颜色
func ResetColor() {
	fmt.Print("\x1b[0m")
}

// SetColor 设置颜色
func SetColor(c Color) {
	fmt.Printf("\x1b[%dm", c)
}