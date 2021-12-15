/*
Project: rainbow fmt.go
Created: 2021/12/15 by Landers
*/

package rainbow

import (
	"bytes"
	"fmt"
)

// 格式化字符串
// 需要前景色 背景色 背景色为nil时不传递

type Colors []Color

// Printf 中的\n不能被正确的添加到颜色格式化串中
// 推荐封装额外的fmt.Print()使用
func Printf(colors Colors,f string, args ...interface{}) {
	if len(colors) == 0 {
		fmt.Printf(f, args...)
		return
	}
	b := new(bytes.Buffer)
	b.WriteString("\x1b[")
	// append styles
	if len(colors) > 1 {
		for _, s := range colors[:len(colors)-1] {
			b.WriteString(s.String())
			b.WriteString(";")
		}
		b.WriteString(colors[len(colors)-1].String())
		b.WriteString("m")
	} else {
		b.WriteString(colors[0].String())
		b.WriteString("m")
	}
	fmt.Printf(
		fmt.Sprintf("%s%%v\x1b[0m", b.String()),
		fmt.Sprintf(f, args...))
}

func Println(colors Colors, args ...interface{}) {
	if len(colors) == 0 {
		fmt.Println(args...)
		return
	}
	b := new(bytes.Buffer)
	b.WriteString("\x1b[")
	// append styles
	if len(colors) > 1 {
		for _, s := range colors[:len(colors)-1] {
			b.WriteString(s.String())
			b.WriteString(";")
		}
		b.WriteString(colors[len(colors)-1].String())
		b.WriteString("m")
	} else {
		b.WriteString(colors[0].String())
		b.WriteString("m")
	}

	fmt.Println(
		fmt.Sprintf("%s%v\x1b[0m", b.String(),
		fmt.Sprint(args...)))
}

func Print(colors Colors, args ...interface{}) {
	if len(colors) == 0 {
		fmt.Print(args...)
		return
	}
	b := new(bytes.Buffer)
	b.WriteString("\x1b[")
	// append styles
	if len(colors) > 1 {
		for _, s := range colors[:len(colors)-1] {
			b.WriteString(s.String())
			b.WriteString(";")
		}
		b.WriteString(colors[len(colors)-1].String())
		b.WriteString("m")
	} else {
		b.WriteString(colors[0].String())
		b.WriteString("m")
	}
	fmt.Print(
		fmt.Sprintf("%s%v\x1b[0m", b.String(),
		fmt.Sprint(args...)))
}
