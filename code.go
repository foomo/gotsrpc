package gotsrpc

import "strings"

type code struct {
	line   string
	lines  []string
	indent int
	tab    string
}

func newCode(tab string) *code {
	return &code{
		line:   "",
		lines:  []string{},
		indent: 0,
		tab:    tab,
	}
}

func (c *code) ind(inc int) *code {
	c.indent += inc
	if c.indent < 0 {
		c.indent = 0
	}
	return c
}

func (c *code) nl() *code {
	c.lines = append(c.lines, strings.Repeat(c.tab, c.indent)+c.line)
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
