package util

import (
	"bytes"
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
	builder *bytes.Buffer
}

func NewText() *Text {
	m := &Text{builder: bytes.NewBuffer(nil)}
	return m
}

func (receiver *Text) WriteString(str string) (int, error) {
	receiver.builder.WriteString(str)
	return len(str), nil
}

func (receiver *Text) WriteByte(str byte) error {
	return receiver.builder.WriteByte(str)
}
func (receiver *Text) String() string {
	return receiver.builder.String()
}

func (receiver *Text) Writef(format string, args ...any) (int, error) {
	return receiver.WriteString(fmt.Sprintf(format, args...))
}
func (receiver *Text) WriteQuote(str string) (int, error) {
	return receiver.Writef("\"%s\"", str)
}

func (receiver *Text) WriteSpace() error {
	return receiver.WriteByte(consts.Space)
}
func (receiver *Text) WriteTab() error {
	return receiver.WriteByte(consts.Tab)
}
func (receiver *Text) WriteStar() error {
	return receiver.WriteByte(consts.Star)
}
func (receiver *Text) WriteEndl() error {
	return receiver.WriteByte(consts.Endl)
}
func (receiver *Text) WriteOpenBrace() error {
	return receiver.WriteByte(consts.OpenBrace)
}
func (receiver *Text) WriteCloseBrace() error {
	return receiver.WriteByte(consts.CloseBrace)
}
func (receiver *Text) WriteTagSign() error {
	return receiver.WriteByte(consts.TagSign)
}

func (receiver *Text) WriteOpenParen() error {
	return receiver.WriteByte(consts.OpenParen)
}
func (receiver *Text) WriteCloseParen() error {
	return receiver.WriteByte(consts.CloseParen)
}
func (receiver *Text) WriteComma() error {
	return receiver.WriteByte(consts.Comma)
}
func (receiver *Text) WriteWithQuote(str string) (int, error) {
	return receiver.Writef("\"%s\"", str)
}
func (receiver *Text) WriteStringWithEndl(str string) (int, error) {
	writeString, err := receiver.WriteString(str)
	if err != nil {
		return 0, err
	}
	return writeString + 1, receiver.WriteEndl()
}
