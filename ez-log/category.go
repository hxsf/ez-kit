package ez_log

import (
	"bytes"

	"github.com/hxsf/ez-kit/ez-buffer"
)

type Category struct {
	Module      string
	Category    string
	SubCategory string
	Filter1     string
	Filter2     string
}

func NewCategory(module string, category string, subCategory string, filter1 string, filter2 string) *Category {
	return &Category{Module: module, Category: category, SubCategory: subCategory, Filter1: filter1, Filter2: filter2}
}

const (
	LeftSep        = byte('[')
	RightSep       = byte(']')
	SpaceSep       = byte(' ')
	AdditionalSize = 3*5 - 1
	Emptycategory = "[] [] [] [] []"
)

func (p *Category) ByteLength() int {
	return len(p.Category) + len(p.SubCategory) + len(p.Module) + len(p.Filter1) + len(p.Filter2) + AdditionalSize
}

func (p *Category) serializeToBuffer(buffer *ez_buffer.Buffer) {
	l := len(p.Category) + len(p.SubCategory) + len(p.Module) + len(p.Filter1) + len(p.Filter2) + AdditionalSize
	buffer.Prealloc(l)
	buffer.AppendByte(LeftSep)
	buffer.AppendString(p.Category)
	buffer.AppendByte(RightSep)
	buffer.AppendByte(SpaceSep)

	buffer.AppendByte(LeftSep)
	buffer.AppendString(p.SubCategory)
	buffer.AppendByte(RightSep)
	buffer.AppendByte(SpaceSep)

	buffer.AppendByte(LeftSep)
	buffer.AppendString(p.Module)
	buffer.AppendByte(RightSep)
	buffer.AppendByte(SpaceSep)

	buffer.AppendByte(LeftSep)
	buffer.AppendString(p.Filter1)
	buffer.AppendByte(RightSep)
	buffer.AppendByte(SpaceSep)

	buffer.AppendByte(LeftSep)
	buffer.AppendString(p.Filter2)
	buffer.AppendByte(RightSep)
}

func (p *Category) Bytes() []byte {
	if p == nil {
		return nil
	}
	buf := &bytes.Buffer{}

	buf.Grow(len(p.Category) + len(p.SubCategory) + len(p.Module) + len(p.Filter1) + len(p.Filter2) + AdditionalSize)

	buf.WriteByte(LeftSep)
	buf.WriteString(p.Category)
	buf.WriteByte(RightSep)
	buf.WriteByte(SpaceSep)

	buf.WriteByte(LeftSep)
	buf.WriteString(p.SubCategory)
	buf.WriteByte(RightSep)
	buf.WriteByte(SpaceSep)

	buf.WriteByte(LeftSep)
	buf.WriteString(p.Module)
	buf.WriteByte(RightSep)
	buf.WriteByte(SpaceSep)

	buf.WriteByte(LeftSep)
	buf.WriteString(p.Filter1)
	buf.WriteByte(RightSep)
	buf.WriteByte(SpaceSep)

	buf.WriteByte(LeftSep)
	buf.WriteString(p.Filter2)
	buf.WriteByte(RightSep)

	return buf.Bytes()
}
