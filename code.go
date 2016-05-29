package gotsrpc

import "strings"

type code struct {
	line   string
	lines  []string
	indent int
}

func newCode() *code {
	return &code{
		line:   "",
		lines:  []string{},
		indent: 0,
	}
}

func (c *code) ind(inc int) *code {
	c.indent += inc
	return c
}

func (c *code) nl() *code {
	c.lines = append(c.lines, strings.Repeat("    ", c.indent)+c.line)
	c.line = ""
	return c
}
func (c *code) l(line string) *code {
	c.app(line).nl()
	return c
}

func (c *code) app(str string) *code {
	c.line += str
	return c
}

func (c *code) string() string {
	return strings.Join(c.lines, "\n")
}
