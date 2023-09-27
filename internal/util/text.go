package util

import (
	"fmt"
	"strings"

	"github.com/wsy998/ast/internal/consts"
)

func UnwrapQuote(str string) string {
	s := []string{"`", `"`}
	sl := ""
	for _, s2 := range s {
		if strings.HasPrefix(str, s2) && strings.HasSuffix(str, s2) {
			sl = str[1 : len(str)-1]
		}
	}
	return sl
}
func EmptyString(str string) bool {
	return str == consts.Empty && len(str) == 0
}
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterLower(s[0]) {
		return string(s[0]-32) + s[1:]
	}
	return s
}
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

type Text struct {
	strings.Builder
}

func (receiver *Text) Writef(format string, args ...any) {
	receiver.WriteString(fmt.Sprintf(format, args...))
}
func (receiver *Text) WriteQuote(str string) {
	receiver.Writef("\"%s\"", str)
}

func (receiver *Text) WriteSpace() {
	receiver.WriteByte(consts.Space)
}
func (receiver *Text) WriteTab() {
	receiver.WriteByte(consts.Tab)
}
func (receiver *Text) WriteStar() {
	receiver.WriteByte(consts.Star)
}
func (receiver *Text) WriteEndl() {
	receiver.WriteByte(consts.Endl)
}
func (receiver *Text) WriteOpenBrace() {
	receiver.WriteByte(consts.OpenBrace)
}
func (receiver *Text) WriteCloseBrace() {
	receiver.WriteByte(consts.CloseBrace)
}
func (receiver *Text) WriteTagSign() {
	receiver.WriteByte(consts.TagSign)
}

func (receiver *Text) WriteOpenParen() {
	receiver.WriteByte(consts.OpenParen)
}
func (receiver *Text) WriteCloseParen() {
	receiver.WriteByte(consts.CloseParen)
}
func (receiver *Text) WriteComma() {
	receiver.WriteByte(consts.Comma)
}
func (receiver *Text) WriteWithQuote(str string) {
	receiver.Writef("\"%s\"", str)
}
func (receiver *Text) WriteStringWithEndl(str string) {
	receiver.WriteString(str)
	receiver.WriteEndl()
}
